# go-mastodon-api

Opinionated and minimalist Go package for working with the Mastodon API.

_This is work in progress. Depending on what you are trying to do you may have more luck with [mattn/go-mastodon](https://github.com/mattn/go-mastodon)._

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

## See also

* https://gocloud.dev/runtimevar
* https://github.com/sfomuseum/runtimevar
* https://github.com/mattn/go-mastodon/
* https://docs.joinmastodon.org/api/