package util

import (
	"flag"
	"os"
)

//Flags represents the flags passed to Donorfide
type Flags struct {
	ClientDebug          bool
	APITester            bool
	DisableStripeWebhook bool
}

func ParseFlags() Flags {
	clientDebug := flag.Bool("client-debug", false, "Enables the client debug mode. Loads the client from the filesystem rather than the embedded client")
	apiTester := flag.Bool("enable-tester", false, "Enables the API tester.")
	disableStripeWebhook := flag.Bool("disable-stripe-webhook", false, "Disables the creating and destroying of stripe webhooks.")

	flag.Parse()

	return Flags{
		ClientDebug:          *clientDebug || os.Getenv("DONORFIDE_CLIENT_DEBUG") == "1",
		APITester:            *apiTester || os.Getenv("DONORFIDE_API_TESTER") == "1",
		DisableStripeWebhook: *disableStripeWebhook || os.Getenv("STRIPE_DISABLE_WEBHOOK") == "1",
	}
}
