package post

import (
	"flag"
	"fmt"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
	"os"
)

// A valid gocloud.dev/runtimevar URI that resolves to a valid aaronland/go-mastodon-api/client URI.
var client_runtimevar_uri string

// The body of the post.
var status string

// The visibility of the post.
var visibility string

// One or paths to local files to append to the post.
var media multi.MultiString

// DefaultFlagSet returns a `flag.FlagSet` instance configured with flags for running the 'post' application.
func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("mastodon")

	fs.StringVar(&client_runtimevar_uri, "client-uri", "", "A valid gocloud.dev/runtimevar URI that resolves to a valid aaronland/go-mastodon-api/client URI.")
	fs.StringVar(&status, "status", "", "The body of the post.")
	fs.StringVar(&visibility, "public", "", "The visibility of the post.")

	fs.Var(&media, "media", "One or paths to local files to append to the post.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Post a message, with zero or more media attachments, to a Mastodon API endpoint.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
	}

	return fs
}
