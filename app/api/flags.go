package api

import (
	"flag"
	"fmt"
	"os"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

// A valid gocloud.dev/runtimevar URI that resolves to a valid aaronland/go-mastodon-api/client URI.
var client_runtimevar_uri string

// Zero or more {KEY}={VALUE} API parameter pairs to include with the API request.
var params multi.KeyValueString

// The HTTP method to issue for the API method.
var http_method string

// A valid Mastodon API endpoint.
var api_method string

var verbose bool

// DefaultFlagSet returns a `flag.FlagSet` instance configured with flags for running the 'api' application.
func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("mastodon")

	fs.StringVar(&client_runtimevar_uri, "client-uri", "", "A valid gocloud.dev/runtimevar URI that resolves to a valid aaronland/go-mastodon-api/client URI.")

	fs.StringVar(&http_method, "http-method", "GET", "The HTTP method to issue for the API method.")
	fs.StringVar(&api_method, "api-method", "", "A valid Mastodon API endpoint.")
	fs.Var(&params, "param", "Zero or more {KEY}={VALUE} API parameter pairs to include with the API request.")

	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Perform an API request against an Mastodon API endpoint.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
	}

	return fs
}
