package requests

type StoreShopRequest struct {
	Name string `form:"name" binding:"required"`
}
