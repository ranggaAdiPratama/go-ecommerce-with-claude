package requests

type LoginRequest struct {
	Username string `json:"username" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required,min=8" validate:"required,min=8"`
}

type RefreshTokenRequest struct {
	Token string `json:"token" binding:"required" validate:"required"`
}
