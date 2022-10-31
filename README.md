# go-mastodon-api

Minimalist and opinionated and minimalist Go package for working with the Mastodon API.

_This is work in progress. When finished its design will be similar to the design of the [aaronland/go-flickr-api](https://github.com/aaronland/go-flickr-api#design) package. Depending on what you are trying to do you may have more luck with [mattn/go-mastodon](https://github.com/mattn/go-mastodon)._

## Documentation

Documentation is incomplete at this time.

## Tools

```
$> make cli
go build -mod vendor -o bin/post cmd/post/main.go
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

## See also

* https://gocloud.dev/runtimevar
* https://github.com/sfomuseum/runtimevar
* https://github.com/mattn/go-mastodon/
* https://docs.joinmastodon.org/api/
