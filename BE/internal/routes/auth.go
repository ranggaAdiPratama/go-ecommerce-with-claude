package routes

import (
	"ranggaAdiPratama/go-with-claude/internal/handlers"
	"ranggaAdiPratama/go-with-claude/internal/middleware"
	"ranggaAdiPratama/go-with-claude/internal/utils"

	"github.com/gin-gonic/gin"
)

func authRoute(r *gin.Engine, h *handlers.AuthHandler, p *utils.PasetoMaker) {
	auth := r.Group("/api/auth")

	auth.POST("/login", h.Login)
	auth.POST("/logout", middleware.AuthMiddleware(p), h.Logout)
	auth.POST("/refresh-token", h.RefreshToken)
	auth.POST("/register", h.Register)

}
