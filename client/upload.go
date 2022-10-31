package client

import (
	"context"
	"fmt"
	"github.com/aaronland/go-mastodon-api/response"
	"os"
)

type uploadResponse struct {
	Index   int
	MediaId string
}

// Upload is a helper method to use 'cl' to upload 'media' to a Mastodon server. It returns the media IDs
// of the the files that were uploaded. If you need to post a message with more options you should
// use the `Client.ExecuteMethod` method instead.
func Upload(ctx context.Context, cl Client, media ...string) ([]string, error) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	count_media := len(media)
	media_ids := make([]string, count_media)

	done_ch := make(chan bool)
	err_ch := make(chan error)
	rsp_ch := make(chan uploadResponse)

	for idx, path := range media {

		go func(ctx context.Context, idx int, path string) {

			defer func() {
				done_ch <- true
			}()

			// To do: Update to read from gocloud.dev/blob.Bucket
			r, err := os.Open(path)

			if err != nil {
				err_ch <- fmt.Errorf("Failed to open path, %w", err)
				return
			}

			defer r.Close()

			rsp, err := cl.UploadMedia(ctx, r, nil)

			if err != nil {
				err_ch <- fmt.Errorf("Failed to upload media, %w", err)
				return
			}

			media_id, err := response.Id(ctx, rsp)

			if err != nil {
				err_ch <- fmt.Errorf("Failed to derive media ID from upload, %w", err)
				return
			}

			rsp_ch <- uploadResponse{Index: idx, MediaId: media_id}

		}(ctx, idx, path)
	}

	remaining := count_media

	for remaining > 0 {
		select {
		case <-ctx.Done():
			return media_ids, nil
		case <-done_ch:
			remaining -= 1
		case err := <-err_ch:
			return nil, err
		case rsp := <-rsp_ch:
			media_ids[rsp.Index] = rsp.MediaId
		}
	}

	return media_ids, nil
}
