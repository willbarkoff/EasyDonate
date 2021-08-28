package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"github.com/stripe/stripe-go/webhook"
	"github.com/stripe/stripe-go/webhookendpoint"
	"github.com/willbarkoff/donorfide/donorfide/database"
	"github.com/willbarkoff/donorfide/donorfide/logging"
)

type paymentIntentResponse struct {
	Status       string `json:"status"`
	ClientSecret string `json:"client_secret"`
}

var webhookEndpoint *stripe.WebhookEndpoint

func setupDonationEndpoints(r *mux.Router) {
	setupStripe()

	r.HandleFunc("/generatePaymentToken", generatePaymentToken).Methods(POST)
	r.HandleFunc("/stripe/webhook", stripeWebhook).Methods(POST)
	r.HandleFunc("/list", listDonations).Methods(GET)

	logging.Logger.Info().Msg("Loaded donation endpoints.")
}

func setupStripe() {
	stripe.Key = database.GetPref(db, "stripeSK")
	stripe.DefaultLeveledLogger = logging.StripeLogger{
		Level: stripe.LevelWarn,
	}

	if flags.DisableStripeWebhook {
		logging.Logger.Info().Msg("Stripe webhook setup disabled. Skipping webhook creation...")
		return
	}
	webhookParams := &stripe.WebhookEndpointParams{
		EnabledEvents: []*string{
			stripe.String("charge.succeeded"),
		},
		URL: stripe.String(database.GetPref(db, "donationPage") + "/api/donate/stripe/webhook"),
	}
	var err error
	webhookEndpoint, err = webhookendpoint.New(webhookParams)
	if err != nil {
		logging.FatalMsg(err, "An error occurred setting up the stripe webhook.")
	} else {
		logging.Logger.Info().Msg("Created stripe webhook")
	}
}

// CleanupStripe removes the stripe webhooks.
func CleanupStripe() {
	if flags.DisableStripeWebhook {
		logging.Logger.Info().Msg("Stripe webhook setup disabled. Skipping webhook deletion...")
		return
	}

	_, err := webhookendpoint.Del(webhookEndpoint.ID, nil)
	if err != nil {
		logging.Logger.Err(err).Msg("Couldn't remove the stripe webhook. Please do it via the stripe dashboard")
	} else {
		logging.Logger.Info().Msg("Removed the stripe webhook.")
	}
}

func generatePaymentToken(w http.ResponseWriter, r *http.Request) {
	if !paramsOk(w, r, "amount") {
		return
	}

	amount, err := strconv.Atoi(r.FormValue("amount"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, statusInvalidParams)
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(int64(amount)),
		Currency:           stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		Description:        stripe.String(database.GetPref(db, "chargeDescription")),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, statusInternalServerError)
		logging.Logger.Err(err).Msg("generating payment intent")
		return
	}

	writeJSON(w, http.StatusOK, okWithData(paymentIntentResponse{
		Status:       "ok",
		ClientSecret: pi.ClientSecret,
	}))
}

func stripeWebhook(w http.ResponseWriter, r *http.Request) {
	const MaxBodyBytes = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logging.Logger.Err(err).Msg("Error reading request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var endpointSecret string
	if webhookEndpoint != nil {
		endpointSecret = webhookEndpoint.Secret
	} else {
		endpointSecret = os.Getenv("DONORFIDE_STRIPE_WEBHOOK_SECRET")
	}

	if endpointSecret == "" {
		logging.Logger.Warn().Msg("The Stripe webhook secret is undefined. Webhook verification will fail.")
	}

	signature := r.Header.Get("Stripe-Signature")
	event, err := webhook.ConstructEvent(payload, signature, endpointSecret)
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, statusUnauthorized)
		return
	}

	switch event.Type {
	case "charge.succeeded":
		var pi stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &pi)
		if err != nil {
			logging.Logger.Warn().Err(err).Str("event.Type", event.Type).Msg("Unable to unmarshall paymentintent")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tx := db.Create(&database.Donation{
			Email:         pi.ReceiptEmail,
			PaymentIntent: pi.ID,
			Currency:      pi.Currency,
			Amount:        pi.Amount,
			Status:        string(pi.Status),
		})
		tx.Commit()

	default:
		logging.Logger.Warn().Str("event.Type", event.Type).Msg("Unhandled event type received from Stripe API.")
	}

	writeJSON(w, http.StatusOK, statusOK)
}

func listDonations(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserId(r, w)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, statusInternalServerError)
		logging.Logger.Err(err).Msg("Getting user ID to list donations.")
		return
	}

	if userID == -1 {
		writeJSON(w, http.StatusUnauthorized, statusLoggedOut)
		return
	}

	userInfo := database.GetUserInfo(db, userID)
	if (userInfo == database.User{}) {
		writeJSON(w, http.StatusUnauthorized, statusUserMissing)
		logging.Logger.Warn().Int("user", userID).Msg("The given user is missing from the database.")
		return
	}

	if userInfo.Level < 1 {
		writeJSON(w, http.StatusUnauthorized, statusUnauthorized)
		return
	}

	var donations []database.Donation
	db.Find(&donations)

	writeJSON(w, http.StatusOK, okWithData(donations))
}
