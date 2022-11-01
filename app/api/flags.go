package api

import (
	"flag"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

var client_runtimevar_uri string

var params multi.KeyValueString

var http_method string

var api_method string

var paginated bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("mastodon")

	fs.StringVar(&client_runtimevar_uri, "client-uri", "", "A valid gocloud.dev/runtimevar URI that resolves to a valid aaronland/go-mastodon-api/client URI.")

	fs.StringVar(&http_method, "http-method", "GET", "The HTTP method to issue for the API method.")
	fs.StringVar(&api_method, "api-method", "", "A valid Mastodon API endpoint.")
	fs.Var(&params, "param", "Zero or more {KEY}={VALUE} API parameter pairs to include with the API request.")

	fs.BoolVar(&paginated, "paginated", false, "Automatically paginate (and iterate through) all results.")
	return fs
}
