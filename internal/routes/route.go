package routes

import (
	"net/http"
	"ranggaAdiPratama/go-with-claude/internal/config"
	"ranggaAdiPratama/go-with-claude/internal/database"
	"ranggaAdiPratama/go-with-claude/internal/handlers"
	"ranggaAdiPratama/go-with-claude/internal/middleware"
	"ranggaAdiPratama/go-with-claude/internal/responses"
	"ranggaAdiPratama/go-with-claude/internal/service"
	"ranggaAdiPratama/go-with-claude/internal/utils"

	"github.com/gin-gonic/gin"
)

func Index(r *gin.Engine, s *database.Store, p *utils.PasetoMaker, c *config.Config) {
	authService := service.NewAuthService(s, p, c)
	userService := service.NewUserService(s)

	authHandler := handlers.NewAuthHandler(authService, userService)
	userHandler := handlers.NewUserHandler(userService)

	r.GET("/", IndexRoute)
	r.GET("/health", HealthRoute)

	auth := r.Group("/api/auth")
	users := r.Group("/api/users")

	auth.POST("/login", authHandler.Login)
	auth.POST("/logout", middleware.AuthMiddleware(p), authHandler.Logout)
	auth.POST("/refresh-token", authHandler.RefreshToken)
	auth.POST("/register", authHandler.Register)

	users.GET("", userHandler.Index)
	users.GET("/:id", userHandler.Show)
	users.POST("", userHandler.Store)
	users.PUT("/:id", userHandler.Update)
	users.DELETE("/:id", userHandler.Destroy)
}

func IndexRoute(c *gin.Context) {
	c.JSON(http.StatusOK, responses.Response{
		MetaData: responses.MetaDataResponse{
			Code:    http.StatusOK,
			Message: "Bijil",
		},
		Data: gin.H{
			"status":  "ok",
			"message": "Created by Rangga Adi Pratama",
		},
	})
}

func HealthRoute(c *gin.Context) {
	c.JSON(http.StatusOK, responses.Response{
		MetaData: responses.MetaDataResponse{
			Code:    http.StatusOK,
			Message: "Service is healthy",
		},
		Data: gin.H{
			"status":  "ok",
			"message": "Service is healthy",
		},
	})
}
