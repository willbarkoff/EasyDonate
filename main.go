package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/BurntSushi/toml"
	"github.com/labstack/echo"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/checkout/session"
	"github.com/stripe/stripe-go/plan"
)

type config struct {
	StripeKey             string `toml:"stripeKey"`
	StripePublicKey       string `toml:"stripePublicKey"`
	ItemName              string `toml:"itemName"`
	BaseURL               string `toml:"baseURL"`
	Name                  string `toml:"name"`
	Disclaimer            string `toml:"disclaimer"`
	SuccessMessage        string `toml:"successMessage"`
	FailureMessage        string `toml:"failureMessage"`
	AllowRecur            bool   `toml:"allowRecur"`
	StopRecurInstructions string `toml:"stopRecurInstructions"`
}

type errorResponse struct {
	Status string `json:"status"`
	Err    string `json:"error"`
}

type sessionIDResponse struct {
	Status    string `json:"status"`
	SessionID string `json:"sessionID"`
}

// Template used for Echo
type Template struct {
	templates *template.Template
}

// Render renders a template
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	var cfg config

	if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
		panic(err)
	}

	stripe.Key = cfg.StripeKey

	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("client/*.html")),
	}
	e.Renderer = t

	e.Static("/", "client")

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", cfg)
	})

	e.GET("/cancel", func(c echo.Context) error {
		return c.Render(http.StatusOK, "cancel.html", cfg)
	})

	e.GET("/success", func(c echo.Context) error {
		return c.Render(http.StatusOK, "success.html", cfg)
	})

	if cfg.AllowRecur {
		e.GET("/stoprecur", func(c echo.Context) error {
			return c.Render(http.StatusOK, "stoprecur.html", cfg)
		})

		e.POST("/recur", func(c echo.Context) error {
			val, err := strconv.Atoi(c.FormValue("amount"))
			if err != nil {
				return c.JSON(http.StatusBadRequest, errorResponse{"error", "invalid params"})
			}

			planParams := &stripe.PlanParams{
				Amount:   stripe.Int64(int64(val)),
				Interval: stripe.String("month"),
				Product: &stripe.PlanProductParams{
					Name: stripe.String("monthly donation"),
				},
				Currency: stripe.String(string(stripe.CurrencyUSD)),
			}
			p, err := plan.New(planParams)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, errorResponse{"error", "invalid params"})
			}

			params := &stripe.CheckoutSessionParams{
				PaymentMethodTypes: stripe.StringSlice([]string{
					"card",
				}),
				SubscriptionData: &stripe.CheckoutSessionSubscriptionDataParams{
					Items: []*stripe.CheckoutSessionSubscriptionDataItemsParams{
						&stripe.CheckoutSessionSubscriptionDataItemsParams{
							Plan: stripe.String(p.ID),
						},
					},
				},
				SuccessURL: stripe.String(cfg.BaseURL + "/success"),
				CancelURL:  stripe.String(cfg.BaseURL + "/cancel"),
			}

			session, err := session.New(params)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, errorResponse{"error", "invalid params"})
			}

			return c.JSON(http.StatusOK, sessionIDResponse{"ok", session.ID})
		})
	}

	e.POST("/donate", func(c echo.Context) error {
		val, err := strconv.Atoi(c.FormValue("amount"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, errorResponse{"error", "invalid params"})
		}

		params := &stripe.CheckoutSessionParams{
			PaymentMethodTypes: stripe.StringSlice([]string{
				"card",
			}),
			LineItems: []*stripe.CheckoutSessionLineItemParams{
				&stripe.CheckoutSessionLineItemParams{
					Name:     stripe.String(cfg.ItemName),
					Amount:   stripe.Int64(int64(val)),
					Currency: stripe.String(string(stripe.CurrencyUSD)),
					Quantity: stripe.Int64(1),
				},
			},
			SuccessURL: stripe.String(cfg.BaseURL + "/success"),
			CancelURL:  stripe.String(cfg.BaseURL + "/cancel"),
		}

		session, err := session.New(params)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, errorResponse{"error", "internal_server_error"})
		}

		return c.JSON(http.StatusOK, sessionIDResponse{"ok", session.ID})
	})

	log.Fatalln(e.Start(":8080"))
}
