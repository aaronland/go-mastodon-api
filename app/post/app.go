package post

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/go-mastodon-api/app"
	"github.com/aaronland/go-mastodon-api/client"
	"github.com/sfomuseum/go-flags/flagset"
	"log"
	"strings"
)

func Run(ctx context.Context, logger *log.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *log.Logger) error {

	flagset.Parse(fs)

	cl, err := app.NewClient(ctx, client_runtimevar_uri, logger)

	if err != nil {
		return fmt.Errorf("Failed to create new client, %w", err)
	}

	status_id, media_ids, err := client.Post(ctx, cl, status, visibility, media...)

	if err != nil {
		return fmt.Errorf("Failed to post message, %w", err)
	}

	fmt.Printf("%s %s\n", status_id, strings.Join(media_ids, ","))
	return nil
}
