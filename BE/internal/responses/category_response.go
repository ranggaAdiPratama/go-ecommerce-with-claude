package responses

import "github.com/google/uuid"

type CategoryPaginatedResponse struct {
	Data       []CategoryResponse `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

type CategoryResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Slug      string    `json:"slug"`
	IsActive  bool      `json:"is_active"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type CategoryResponseShort struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Icon string    `json:"icon"`
	Slug string    `json:"slug"`
}
