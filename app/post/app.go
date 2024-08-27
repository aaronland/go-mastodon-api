package post

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"strings"

	"github.com/aaronland/go-mastodon-api/app"
	"github.com/aaronland/go-mastodon-api/client"
	"github.com/sfomuseum/go-flags/flagset"
)

// Run will execute the 'post' commandline application with default flags.
func Run(ctx context.Context) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

// Run will execute the 'post' commandline application with 'fs'.
func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "MASTODON")

	if err != nil {
		return fmt.Errorf("Failed to set flags from environment variables, %w", err)
	}

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	cl, err := app.NewClient(ctx, client_runtimevar_uri)

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
