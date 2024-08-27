package api

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net/url"
	"os"

	"github.com/aaronland/go-mastodon-api/app"
	"github.com/sfomuseum/go-flags/flagset"
)

// Run will execute the 'api' commandline application with default flags.
func Run(ctx context.Context) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

// Run will execute the 'api' commandline application with 'fs'.
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

	args := &url.Values{}

	for _, kv := range params {
		args.Add(kv.Key(), kv.Value().(string))
	}

	rsp, err := cl.ExecuteMethod(ctx, http_method, api_method, args)

	if err != nil {
		return fmt.Errorf("Failed to execute method, %w", err)
	}

	_, err = io.Copy(os.Stdout, rsp)

	if err != nil {
		return fmt.Errorf("Failed to emit response, %w", err)
	}

	return nil
}
