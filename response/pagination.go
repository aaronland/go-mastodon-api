package response

import (
	"context"
	"fmt"
	"io"
)

// Pagination is a struct containing pagination metrics for a given API response.
type Pagination struct {
	// The current page of results for an API request.
	Page int `json:"page"`
	// The total number of pages of results for an API request.
	Pages int `json:"pages"`
	// The number of results, per page, for an API request.
	PerPage int `json:"perpage"`
	// The total number of results, across all pages, for an API request.
	Total int `json:total"`
}

// Still WIP but include to make the client.ExecuteMethodPaginated method compile.

func DerivePagination(ctx context.Context, r io.ReadSeeker) (*Pagination, error) {
	return nil, fmt.Errorf("Not implemented")
}
