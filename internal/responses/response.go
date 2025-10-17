package responses

import "github.com/google/uuid"

type Response struct {
	MetaData MetaDataResponse `json:"meta_data"`
	Data     any              `json:"data"`
}

type MetaDataResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
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
