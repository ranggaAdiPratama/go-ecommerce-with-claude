package routes

import (
	"ranggaAdiPratama/go-with-claude/internal/handlers"
	"ranggaAdiPratama/go-with-claude/internal/middleware"
	"ranggaAdiPratama/go-with-claude/internal/utils"

	"github.com/gin-gonic/gin"
)

func settingRoute(r *gin.Engine, h *handlers.SettingHandler, p *utils.PasetoMaker) {
	setting := r.Group("/api/settings")

	setting.GET("", h.Index)
	setting.POST("", middleware.AuthMiddleware(p), middleware.RequireRole("admin"), h.StoreOrUpdate)
}
