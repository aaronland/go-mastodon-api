package main

import (
	"context"
	"log"

	app "github.com/aaronland/go-mastodon-api/v2/app/api"
)

func main() {

	ctx := context.Background()
	err := app.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run application, %v", err)
	}
}
