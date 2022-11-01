// Package client provides a common interface for accessing the Mastodon API.
package client

import (
	"context"
	"fmt"
	"github.com/aaronland/go-roster"
	"io"
	"log"
	"net/url"
	"sort"
	"strings"
)

// Client is the interface that defines common methods for all Mastodon API Client implementations.
// Currently there is only a single implementation that calls the Mastodon API using the OAuth2
// authentication and authorization scheme but it is assumed that eventually there will be others.
type Client interface {
	// Execute a Mastodon API method.
	ExecuteMethod(context.Context, string, string, *url.Values) (io.ReadSeekCloser, error)
	// Upload an io.Reader instance using the Mastodon API.
	UploadMedia(context.Context, io.Reader, *url.Values) (io.ReadSeekCloser, error)
	// Assign a specific log.Logger instance for logging events.
	SetLogger(context.Context, *log.Logger) error
}

var clients roster.Roster

// The initialization function signature for implementation of the Client interface.
type ClientInitializeFunc func(context.Context, string) (Client, error)

// Ensure that the internal roster.Roster instance has been created successfully.
func ensureClientRoster() error {

	if clients == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		clients = r
	}

	return nil
}

// Register a new URI scheme and ClientInitializeFunc function for a implementation of the Client interface.
func RegisterClient(ctx context.Context, scheme string, f ClientInitializeFunc) error {

	err := ensureClientRoster()

	if err != nil {
		return err
	}

	return clients.Register(ctx, scheme, f)
}

// Return a list of URI schemes for registered implementations of the Client interface.
func Schemes() []string {

	ctx := context.Background()
	schemes := []string{}

	err := ensureClientRoster()

	if err != nil {
		return schemes
	}

	for _, dr := range clients.Drivers(ctx) {
		scheme := fmt.Sprintf("%s://", strings.ToLower(dr))
		schemes = append(schemes, scheme)
	}

	sort.Strings(schemes)
	return schemes
}

// Create a new instance of the Client interface. Client instances are created by
// passing in a context.Context instance and a URI string. The form and substance of
// URI strings are specific to their implementations. For example to create a OAuth2Client
// you would write:
// cl, err := client.NewClient(ctx, "oauth2://:{ACCESS_TOKEN}@{MASTODON_HOST}")
func NewClient(ctx context.Context, uri string) (Client, error) {

	// To account for things that might be gocloud.dev/runtimevar-encoded
	// in a file using editors that automatically add newlines (thanks, Emacs)

	uri = strings.TrimSpace(uri)

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := clients.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	f := i.(ClientInitializeFunc)
	return f(ctx, uri)
}
