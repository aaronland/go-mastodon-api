package client

import (
	"context"
	"fmt"
	"github.com/aaronland/go-mastodon-api/response"
	"io"
	"log"
	"net/url"
	"strconv"
)

type Client interface {
	ExecuteMethod(context.Context, string, string, *url.Values) (io.ReadSeekCloser, error)
	UploadMedia(context.Context, io.Reader, *url.Values) (io.ReadSeekCloser, error)
	SetLogger(context.Context, *log.Logger) error
}

// ExecuteMethodPaginatedCallback is the interface for callback functions passed to the
// ExecuteMethodPaginatedWithClient method.
type ExecuteMethodPaginatedCallback func(context.Context, io.ReadSeekCloser, error) error

func ExecuteMethodPaginatedWithClient(ctx context.Context, cl Client, http_method string, api_method string, args *url.Values, cb ExecuteMethodPaginatedCallback) error {

	page := 1
	pages := -1

	if args.Get("page") == "" {
		args.Set("page", strconv.Itoa(page))
	} else {

		p, err := strconv.Atoi(args.Get("page"))

		if err != nil {
			return fmt.Errorf("Invalid page number '%s', %v", args.Get("page"), err)
		}

		page = p
	}

	for {

		select {
		case <-ctx.Done():
			return nil
		default:
			// pass
		}

		fh, err := cl.ExecuteMethod(ctx, http_method, api_method, args)

		err = cb(ctx, fh, err)

		if err != nil {
			return err
		}

		_, err = fh.Seek(0, 0)

		if err != nil {
			return fmt.Errorf("Failed to rewind response, %v", err)
		}

		if pages == -1 {

			pagination, err := response.DerivePagination(ctx, fh)

			if err != nil {
				return err
			}

			pages = pagination.Pages
		}

		page += 1

		if page <= pages {
			args.Set("page", strconv.Itoa(page))
		} else {
			break
		}
	}

	return nil
}
