package routes

import (
	"ranggaAdiPratama/go-with-claude/internal/handlers"
	"ranggaAdiPratama/go-with-claude/internal/middleware"
	"ranggaAdiPratama/go-with-claude/internal/utils"

	"github.com/gin-gonic/gin"
)

func shopRoute(r *gin.Engine, h *handlers.ShopHandler, p *utils.PasetoMaker) {
	myShop := r.Group("/api/my-shop").Use(middleware.AuthMiddleware(p))
	shops := r.Group("/api/shops")

	myShop.POST("", h.Store)
	myShop.PUT("", middleware.RequireRole("user"), h.UpdatePersonal)

	shops.GET("", h.Index)
}
