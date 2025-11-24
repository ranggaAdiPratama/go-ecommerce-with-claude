package routes

import (
	"ranggaAdiPratama/go-with-claude/internal/handlers"
	"ranggaAdiPratama/go-with-claude/internal/middleware"
	"ranggaAdiPratama/go-with-claude/internal/utils"

	"github.com/gin-gonic/gin"
)

func categoryRoute(r *gin.Engine, h *handlers.CategoryHandler, p *utils.PasetoMaker) {
	categories := r.Group("/api/categories")

	categories.GET("", h.Index)
	categories.GET("/data", middleware.AuthMiddleware(p), middleware.RequireRole("admin"), h.DataTable)
	categories.GET("/:slug", h.GetBySlug)
	categories.POST("", middleware.AuthMiddleware(p), middleware.RequireRole("admin"), h.Store)
}
