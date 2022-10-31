package api

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/go-mastodon-api/client"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/runtimevar"
	"io"
	"log"
	"net/url"
	"os"
)

func Run(ctx context.Context, logger *log.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *log.Logger) error {

	flagset.Parse(fs)

	client_uri, err := runtimevar.StringVar(ctx, client_runtimevar_uri)

	if err != nil {
		return fmt.Errorf("Failed to derive client uri, %v", err)
	}

	cl, err := client.NewClient(ctx, client_uri)

	if err != nil {
		return fmt.Errorf("Failed to create client, %v", err)
	}

	err = cl.SetLogger(ctx, logger)

	if err != nil {
		return fmt.Errorf("Failed to set logger, %w", err)
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
