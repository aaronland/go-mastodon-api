# go-mastodon-api

Minimalist and opinionated Go package for working with the Mastodon API.

_This is work in progress. When finished its design will be similar to the design of the [aaronland/go-flickr-api](https://github.com/aaronland/go-flickr-api#design) package. Depending on what you are trying to do you may have more luck with [mattn/go-mastodon](https://github.com/mattn/go-mastodon)._

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/aaronland/go-mastodon-api.svg)](https://pkg.go.dev/github.com/aaronland/go-mastodon-api)

_Documentation is incomplete at this time._

## Design

The core of this package's approach to the Mastodon API is the `ExecuteMethod` method (which is defined in the `client.Client` interface) whose signature looks like this:

```
ExecuteMethod(context.Context, string, string, *url.Values) (io.ReadSeekCloser, error)
```

In time there may be others, along with helper methods for unmarshaling API responses in to typed responses but the baseline for all operations will remain: Query paramters (`url.Values`) sent over HTTP returning an `io.ReadSeekCloser` instance that is inspected and validated according to the needs and uses of the tools using the Mastodon API.

### Pagination

This package does not provide any methods for dealing with pagination since there don't appear to be any common patterns in the Mastodon API from which pagination methods can be derived. Maybe some day?

## Clients

The `client.Client` interface provides for common methods for accessing the Mastodon API. Currently there is only a single client interface that calls the Mastodon API using the OAuth2 authentication and authorization scheme but it is assumed that eventually there will be others.

Clients are instantiated using a URI-based syntax where the scheme and query parameters map to specific implementation details.

### OAuth2

The OAuth2 `Client` implementation is instantiated using the `oauth2://` scheme. For example:

```
oauth2://:{ACCESS_TOKEN}@{MASTODON_HOST}
```

## Tools

```
> make cli
go build -mod vendor -o bin/post cmd/post/main.go
go build -mod vendor -o bin/api cmd/api/main.go
```

### api

```
$> ./bin/api -h
  -api-method string
    	A valid Mastodon API endpoint.
  -client-uri string
    	A valid gocloud.dev/runtimevar URI that resolves to a valid aaronland/go-mastodon-api/client URI.
  -http-method string
    	The HTTP method to issue for the API method. (default "GET")
  -param value
    	Zero or more {KEY}={VALUE} API parameter pairs to include with the API request.
```

For example:

```
$> ./bin/api \
	-client-uri 'file:///usr/local/mastodon/account-uri' \
	-api-method /api/v2/search \
	-param q=caturday \
	-param limit=1 \
	| jq
	
{
  "accounts": [],
  "statuses": [],
  "hashtags": [
    {
      "name": "cáturday",
      "url": "https://mastodon.cloud/tags/c%C3%A1turday",
      "history": [
        {
          "day": "1667174400",
          "uses": "0",
          "accounts": "0"
        },
        {
          "day": "1667088000",
          "uses": "0",
          "accounts": "0"
        },
	... and so on
```

### post

```
$> ./bin/post -h
  -client-uri string
    	A valid gocloud.dev/runtimevar URI that resolves to a valid aaronland/go-mastodon-api/client URI.
  -media value
    	One or paths to local files to append to the post.
  -public string
    	The visibility of the post.
  -status string
    	The body of the post.
```

For example:

```
$> ./bin/post \
	-client-uri 'file:///usr/local/mastodon/account-uri' \
	-status '[testing] test11' \
	-media /usr/local/example.jpg

# post ID followed by media IDs
109259899346500374 109259899212388461
```

## Client URIs

Client URIs take the form of:

```
{SCHEME}://:{ACCESS_TOKEN}@{MASTODON_HOST}
```

For example:

```
oauth2://:S33KRET@mastodon.example
```

Notes, as of this writing:

* There is only one valid scheme for client URIs: `oauth2`.
* There are no methods for doing the OAuth2 access token flow. It is assumed that you have created one by some other means (for example, by generating a new application and access token in the Mastodon web interface). 

## Runtimevar URIs

Under the hood this package uses the [sfomuseum/runtimevar](https://github.com/sfomuseum/runtimevar) package. By default it imports the following [gocloud.dev/runtimevar](https://gocloud.dev/howto/runtimevar) implementations:

* [Constant](https://gocloud.dev/howto/runtimevar/#constant)
* [Local](https://gocloud.dev/howto/runtimevar/#local)
* [AWS Parameter Store](https://gocloud.dev/howto/runtimevar/#awsps)

If you need to import a different implementation you will need to clone the relevant tool and add your custom import statement. For example, to use the `post` tool with [the GCP Secret Manager `runtimevar` implementation](https://gocloud.dev/howto/runtimevar/#gcpsm) you would do:

```
package main

import (
	"context"
	"github.com/aaronland/go-mastodon-api/app/post"
	_ "gocloud.dev/runtimevar/gcpsecretmanager"
	"log"
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	err := post.Run(ctx, logger)

	if err != nil {
		logger.Fatalf("Failed to run application, %v", err)
	}
}
```

## See also

* https://gocloud.dev/runtimevar
* https://github.com/sfomuseum/runtimevar
* https://github.com/mattn/go-mastodon/
* https://docs.joinmastodon.org/api/
