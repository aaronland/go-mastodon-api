package app

import (
	"context"
	"fmt"
	"github.com/aaronland/go-mastodon-api/client"
	"github.com/sfomuseum/runtimevar"
	"log"
)

// NewClient will return a new `client.Client` instance whose URI constructor will be derived from 'client_runtimevar_uri'
// and configured to use 'logger'.
func NewClient(ctx context.Context, client_runtimevar_uri string, logger *log.Logger) (client.Client, error) {

	client_uri, err := runtimevar.StringVar(ctx, client_runtimevar_uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive client uri, %v", err)
	}

	cl, err := client.NewClient(ctx, client_uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create client, %v", err)
	}

	err = cl.SetLogger(ctx, logger)

	if err != nil {
		return nil, fmt.Errorf("Failed to set logger, %w", err)
	}

	return cl, nil
}
