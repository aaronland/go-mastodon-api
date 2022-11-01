package response

import (
	"context"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
)

// Id derives the value of a top-level "id" JSON-encoded property (key) from 'r'.
func Id(ctx context.Context, r io.ReadSeeker) (string, error) {

	defer r.Seek(0, 0)

	body, err := io.ReadAll(r)

	if err != nil {
		return "", fmt.Errorf("Failed to read response, %w", err)
	}

	id_rsp := gjson.GetBytes(body, "id")

	if !id_rsp.Exists() {
		return "", fmt.Errorf("Top-level 'id' property does not exist")
	}

	return id_rsp.String(), nil
}
