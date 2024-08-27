package main

import (
	"context"
	"log"

	"github.com/aaronland/go-mastodon-api/v2/app/post"
)

func main() {

	ctx := context.Background()
	err := post.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run application, %v", err)
	}
}
