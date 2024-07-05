package repository

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type PaginateRequest struct {
	Filters    []Filter
	Search     map[string]string
	Pagination Pagination
}

type PaginatedResult[T any] struct {
	Data      []*T  `json:"data"`
	Total     int64 `json:"total"`
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	PageCount int   `json:"page_count"`
}

type Filter struct {
	Column string
	Value  any
}

type CreateEntityResponse struct {
	ID uint
}
