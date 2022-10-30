package post

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/go-mastodon-api/client"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/runtimevar"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/url"
	"os"
	"strings"
)

func Run(ctx context.Context, logger *log.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *log.Logger) error {

	flagset.Parse(fs)

	client_uri, err := runtimevar.StringVar(ctx, client_runtimevar_uri)

	if err != nil {
		return fmt.Errorf("Failed to derive client uri, %v", err)
	}

	cl, err := client.NewClient(ctx, client_uri)

	if err != nil {
		return fmt.Errorf("Failed to create client, %v", err)
	}

	err = cl.SetLogger(ctx, logger)

	if err != nil {
		return fmt.Errorf("Failed to set logger, %w", err)
	}

	args := &url.Values{}
	args.Set("status", status)
	args.Set("visibility", visibility)

	// START OF handle media

	count_media := len(media)

	var media_ids []string

	if count_media > 0 {

		media_ids = make([]string, count_media)

		for idx, path := range media {

			// To do: Update to read from gocloud.dev/blob.Bucket
			r, err := os.Open(path)

			if err != nil {
				return fmt.Errorf("Failed to open path, %v", err)
			}

			defer r.Close()

			rsp, err := cl.UploadMedia(ctx, r, nil)

			if err != nil {
				return fmt.Errorf("Failed to upload media, %v", err)
			}

			body, err := io.ReadAll(rsp)

			if err != nil {
				return fmt.Errorf("Failed read media upload response, %v", err)
			}

			media_id, err := getId(body)

			if err != nil {
				return fmt.Errorf("Failed to derive media ID from upload, %v", err)
			}

			media_ids[idx] = media_id
			args.Add("media_ids[]", media_id)
		}
	}

	// END OF handle media

	r, err := cl.ExecuteMethod(ctx, "POST", "/api/v1/statuses", args)

	if err != nil {
		return fmt.Errorf("Failed to post message, %v", err)
	}

	body, err := io.ReadAll(r)

	if err != nil {
		return fmt.Errorf("Failed to read response, %v", err)
	}

	status_id, err := getId(body)

	if err != nil {
		return fmt.Errorf("Failed to derive status ID, %v", err)
	}

	fmt.Printf("%s %s\n", status_id, strings.Join(media_ids, ","))
	return nil
}

func getId(body []byte) (string, error) {

	id_rsp := gjson.GetBytes(body, "id")

	if !id_rsp.Exists() {
		return "", fmt.Errorf("Failed to derive ID from post response")
	}

	return id_rsp.String(), nil
}
