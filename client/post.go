package client

import (
	"context"
	"fmt"
	"github.com/aaronland/go-mastodon-api/response"
	"net/url"
)

// Post is a helper method to use 'cl' to post a message to Mastodon with 'post' and 'visibility' and zero or more media files.
// It returns the ID of the post and of the IDs of any media uploads. If you need to post a message with more options you should
// use the `Client.ExecuteMethod` method instead.
func Post(ctx context.Context, cl Client, status string, visibility string, media ...string) (string, []string, error) {

	args := &url.Values{}
	args.Set("status", status)
	args.Set("visibility", visibility)

	count_media := len(media)

	var media_ids []string

	if count_media > 0 {

		m_ids, err := Upload(ctx, cl, media...)

		if err != nil {
			return "", nil, fmt.Errorf("Failed to upload media, %w", err)
		}

		media_ids = m_ids

		for _, media_id := range media_ids {
			args.Add("media_ids[]", media_id)
		}
	}

	r, err := cl.ExecuteMethod(ctx, "POST", "/api/v1/statuses", args)

	if err != nil {
		return "", nil, fmt.Errorf("Failed to post message, %w", err)
	}

	status_id, err := response.Id(ctx, r)

	if err != nil {
		return "", nil, fmt.Errorf("Failed to derive status ID, %w", err)
	}

	return status_id, media_ids, nil
}
