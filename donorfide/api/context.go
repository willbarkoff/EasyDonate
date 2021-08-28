package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/willbarkoff/donorfide/donorfide/database"
	"github.com/willbarkoff/donorfide/donorfide/logging"
)

type orgContext struct {
	StripePK                   string `json:"stripe_pk"`
	Name                       string `json:"name"`
	Site                       string `json:"site"`
	Phone                      string `json:"phone"`
	Email                      string `json:"email"`
	Imprint                    string `json:"imprint"`
	DonatePageTitle            string `json:"donate_page_title"`
	DonatePageSubtitle         string `json:"donate_page_subtitle"`
	DefaultDonationAmount      int    `json:"default_donation_amount"`
	HomePageText               string `json:"home_page_text"`
	CreditPageText             string `json:"credit_page_text"`
	DisableConfetti            bool   `json:"disable_confetti"`
	DonationSuccessRedirect    string `json:"donation_success_redirect"`
	DonationSuccessHeadline    string `json:"donation_success_headline"`
	DonationSuccessSubheadline string `json:"donation_success_subheadline"`
	DonationSuccessText        string `json:"donation_success_text"`
	ReceiptText                string `json:"receipt_text"`
}

func setupContextEndpoints(r *mux.Router) {
	openSessionsTable()

	r.HandleFunc("/org", contextOrg).Methods(GET)

	logging.Logger.Info().Msg("Loaded context endpoints.")
}

func contextOrg(w http.ResponseWriter, r *http.Request) {
	defaultDonationAmount, _ := strconv.Atoi(database.GetPref(db, "defaultDonationAmount"))

	writeJSON(w, http.StatusOK, okWithData(orgContext{
		StripePK:                   database.GetPref(db, "stripePK"),
		Name:                       database.GetPref(db, "orgName"),
		Site:                       database.GetPref(db, "orgSite"),
		Phone:                      database.GetPref(db, "orgPhone"),
		Email:                      database.GetPref(db, "orgEmail"),
		Imprint:                    database.GetPref(db, "orgImprint"),
		DonatePageTitle:            database.GetPref(db, "donatePageTitle"),
		DonatePageSubtitle:         database.GetPref(db, "donatePageSubtitle"),
		DefaultDonationAmount:      defaultDonationAmount,
		HomePageText:               database.GetPref(db, "homePageText"),
		CreditPageText:             database.GetPref(db, "creditPageText"),
		DisableConfetti:            database.GetPref(db, "disableConfetti") == "true",
		DonationSuccessRedirect:    database.GetPref(db, "donationSuccessRedirect"),
		DonationSuccessHeadline:    database.GetPref(db, "donationSuccessHeadline"),
		DonationSuccessSubheadline: database.GetPref(db, "donationSuccessSubheadline"),
		DonationSuccessText:        database.GetPref(db, "donationSuccessText"),
	}))
}
