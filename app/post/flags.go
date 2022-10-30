package post

import (
	"flag"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

var client_runtimevar_uri string
var status string
var visibility string

var media multi.MultiString

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("mastodon")

	fs.StringVar(&client_runtimevar_uri, "client-uri", "", "A valid gocloud.dev/runtimevar URI that resolves to a valid aaronland/go-mastodon-api/client URI.")
	fs.StringVar(&status, "status", "", "The body of the post.")
	fs.StringVar(&visibility, "public", "", "The visibility of the post.")

	fs.Var(&media, "media", "One or paths to local files to append to the post.")

	return fs
}
