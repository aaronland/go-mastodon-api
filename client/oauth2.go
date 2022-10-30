package client

import (
	"bytes"
	"context"
	"fmt"
	"github.com/whosonfirst/go-ioutil"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
)

func init() {

	ctx := context.Background()
	err := RegisterClient(ctx, "oauth2", NewOAuth2Client)

	if err != nil {
		panic(err)
	}
}

type OAuth2Client struct {
	http_client  *http.Client
	api_endpoint *url.URL
	access_token string
	logger       *log.Logger
}

func NewOAuth2Client(ctx context.Context, uri string) (Client, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	mastodon_url := fmt.Sprintf("https://%s", u.Host)

	mastodon_endpoint, err := url.Parse(mastodon_url)

	if err != nil {
		return nil, fmt.Errorf("Invalid Mastodon host, %w", err)
	}

	http_client := &http.Client{}

	logger := log.Default()

	cl := &OAuth2Client{
		http_client:  http_client,
		api_endpoint: mastodon_endpoint,
		logger:       logger,
	}

	token, ok := u.User.Password()

	if ok {
		cl.access_token = token
	}

	return cl, nil
}

func (cl *OAuth2Client) SetLogger(ctx context.Context, logger *log.Logger) error {
	cl.logger = logger
	return nil
}

func (cl *OAuth2Client) ExecuteMethod(ctx context.Context, http_method string, api_method string, args *url.Values) (io.ReadSeekCloser, error) {

	req_endpoint, err := cl.requestEndpoint(ctx, api_method)

	if err != nil {
		return nil, err
	}

	req_endpoint.RawQuery = args.Encode()

	req, err := http.NewRequest(http_method, req_endpoint.String(), nil)

	if err != nil {
		return nil, err
	}

	return cl.call(ctx, req)
}

func (cl *OAuth2Client) UploadMedia(ctx context.Context, r io.Reader, args *url.Values) (io.ReadSeekCloser, error) {

	return cl.upload(ctx, r, args)
}

func (cl *OAuth2Client) upload(ctx context.Context, r io.Reader, args *url.Values) (io.ReadSeekCloser, error) {

	http_method := "POST"
	api_method := "/api/v1/media"

	req_endpoint, err := cl.requestEndpoint(ctx, api_method)

	if err != nil {
		return nil, err
	}

	// This is what I would like to do (see notes in upload.go)

	/*
		boundary, err := randomBoundary()

		if err != nil {
			return nil, err
		}

		r, w := io.Pipe()

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		go func() {

			err := streamUploadBody(ctx, w, "file", boundary, r, args)

			if err != nil {
				log.Printf("Failed to stream upload body for file, %v", err)
				cancel()
			}
		}()
	*/

	// This is what we do instead (womp womp...)

	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)

	file, err := mw.CreateFormFile("file", "upload")

	if err != nil {
		return nil, fmt.Errorf("Failed to create form file, %w", err)
	}

	_, err = io.Copy(file, r)

	if err != nil {
		return nil, fmt.Errorf("Failed to copy media to form file, %w", err)
	}

	err = mw.Close()

	if err != nil {
		return nil, fmt.Errorf("Failed to close form file, %w", err)
	}

	r = bytes.NewReader(buf.Bytes())

	req, err := http.NewRequest(http_method, req_endpoint.String(), r)

	if err != nil {
		return nil, fmt.Errorf("Failed to create upload request, %w", err)
	}

	req.Header.Set("Content-Type", mw.FormDataContentType())

	return cl.call(ctx, req)
}

func (cl *OAuth2Client) call(ctx context.Context, req *http.Request) (io.ReadSeekCloser, error) {

	req = req.WithContext(ctx)

	if cl.access_token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.access_token))
	}

	rsp, err := cl.http_client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Failed to do request, %w", err)
	}

	if rsp.StatusCode != http.StatusOK {
		rsp.Body.Close()
		return nil, fmt.Errorf("API call failed with status '%s'", rsp.Status)
	}

	return ioutil.NewReadSeekCloser(rsp.Body)
}

func (cl *OAuth2Client) requestEndpoint(ctx context.Context, api_method string) (*url.URL, error) {

	req_endpoint, err := url.Parse(cl.api_endpoint.String())

	if err != nil {
		return nil, err
	}

	req_endpoint.Path = api_method
	return req_endpoint, nil
}
