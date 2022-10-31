package main

import (
	"context"
	app "github.com/aaronland/go-mastodon-api/app/api"
	"log"
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	err := app.Run(ctx, logger)

	if err != nil {
		logger.Fatalf("Failed to run application, %v", err)
	}
}
