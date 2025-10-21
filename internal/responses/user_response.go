package responses

import "github.com/google/uuid"

type UserListResponse struct {
	Data       []UserResponse     `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}
