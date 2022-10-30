package client

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"path/filepath"
)

// Most of the code in this file has been copypasted with minor
// updates from https://github.com/masci/flickr/blob/v2/upload.go

// Generate a random multipart boundary string, shamelessly copypasted from the std library
func randomBoundary() (string, error) {

	var buf [30]byte

	_, err := io.ReadFull(rand.Reader, buf[:])
	if err != nil {
		return "", err
	}

	boundary := fmt.Sprintf("%x", buf[:])
	return boundary, nil
}

/*

This does not work and fails with the following error, which I don't completely understand:

		> go run cmd/post/main.go -client-uri 'file:///usr/local/mastodon/orthis-uri' -status '[testing] test6' -media ~/Desktop/000.jpg
		2022/10/30 13:17:46 Failed to stream upload body for 'file', io: read/write on closed pipe
		2022/10/30 13:17:46 Failed to upload media, Post "https://example.com/api/v1/media": http2: Transport: cannot retry err [stream error: stream ID 1; PROTOCOL_ERROR; received from peer] after Request.Body was written; define Request.GetBody to avoid this error
		exit status 1

*/

// Encode the file and request parameters in a multipart body.
// File contents are streamed into the request using an io.Pipe in a separated goroutine
func streamUploadBody(ctx context.Context, body *io.PipeWriter, file_name string, boundary string, fh io.Reader, args *url.Values) error {

	// multipart writer to fill the body
	defer body.Close()

	writer := multipart.NewWriter(body)
	// writer.SetBoundary(boundary)

	part, err := writer.CreateFormFile(file_name, filepath.Base(file_name))

	if err != nil {
		return err
	}

	_, err = io.Copy(part, fh)

	if err != nil {
		return err
	}

	// close the form writer
	return writer.Close()
}
