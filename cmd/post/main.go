package main

import (
	"context"
	"github.com/aaronland/go-mastodon-api/app/post"
	"log"
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	err := post.Run(ctx, logger)

	if err != nil {
		logger.Fatalf("Failed to run application, %v", err)
	}
}
