package requests

type SettingRequest struct {
	Name string `form:"name" binding:"required"`
}
