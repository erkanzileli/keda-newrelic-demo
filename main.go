package main

import (
	"log"
	"os"
	"time"

	"github.com/erkanzileli/nrfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	appName := os.Getenv("NEW_RELIC_APP_NAME")
	licenseKey := os.Getenv("NEW_RELIC_LICENSE_KEY")
	app := fiber.New()
	nr, err := newrelic.NewApplication(newrelic.ConfigEnabled(true), newrelic.ConfigAppName(appName), newrelic.ConfigLicense(licenseKey))
	if err != nil {
		log.Fatal(err)
	}

	// Add the nrfiber middleware before other middlewares or routes
	app.Use(nrfiber.Middleware(nr))

	// Use created transaction to create custom segments
	app.Get("/cart", func(ctx *fiber.Ctx) error {
		txn := nrfiber.FromContext(ctx)
		segment := txn.StartSegment("Price Calculation")
		defer segment.End()

		time.Sleep(2 * time.Second)

		return nil
	})
	app.Listen(":3000")
}
