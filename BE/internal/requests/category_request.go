package requests

type StoreCategoryRequest struct {
	Name     string `json:"name" binding:"required"`
	Icon     string `json:"icon" binding:"required"`
	IsActive string `json:"is_active" binding:"omitempty,oneof=1 0"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
	Icon string `json:"icon" binding:"required"`
}
