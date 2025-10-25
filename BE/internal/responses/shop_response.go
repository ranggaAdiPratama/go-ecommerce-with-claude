package responses

import (
	"github.com/google/uuid"
)

type ShopDisplayResponse struct {
	Data       []ShopResponseShort `json:"data"`
	Pagination PaginationResponse  `json:"pagination"`
}

type ShopResponseForUser struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Logo      string    `json:"logo"`
	Rank      string    `json:"rank"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type ShopResponseShort struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Logo string    `json:"logo"`
	Rank string    `json:"rank"`
}
