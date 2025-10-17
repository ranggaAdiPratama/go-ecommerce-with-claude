package requests

type StoreUserRequest struct {
	Name     string `json:"name" binding:"required" validate:"required"`
	Email    string `json:"email" binding:"required,email" validate:"required,email"`
	Username string `json:"username" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required,min=8" validate:"required,min=8"`
	Role     string `json:"role" binding:"required" validate:"required"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" binding:"required" validate:"required"`
	Email    string `json:"email" binding:"required,email" validate:"required,email"`
	Username string `json:"username" binding:"required" validate:"required"`
	Password string `json:"password" binding:"omitempty,min=8" validate:"min=8"`
	Role     string `json:"role" binding:"omitempty"`
}
