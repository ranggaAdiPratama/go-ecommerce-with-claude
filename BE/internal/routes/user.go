package routes

import (
	"ranggaAdiPratama/go-with-claude/internal/handlers"
	"ranggaAdiPratama/go-with-claude/internal/middleware"
	"ranggaAdiPratama/go-with-claude/internal/utils"

	"github.com/gin-gonic/gin"
)

func userRoute(r *gin.Engine, h *handlers.UserHandler, p *utils.PasetoMaker) {
	users := r.Group("/api/users").Use(middleware.AuthMiddleware(p)).Use(middleware.RequireRole("admin"))

	users.GET("", h.Index)
	users.GET("/:id", h.Show)
	users.POST("", h.Store)
	users.PUT("/:id", h.Update)
	users.DELETE("/:id", h.Destroy)
}
