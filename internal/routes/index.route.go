package routes

import (
	"net/http"
	"ranggaAdiPratama/go-with-claude/internal/database"
	"ranggaAdiPratama/go-with-claude/internal/handlers"
	"ranggaAdiPratama/go-with-claude/internal/responses"
	"ranggaAdiPratama/go-with-claude/internal/service"

	"github.com/gin-gonic/gin"
)

func Index(r *gin.Engine, s *database.Store) {
	userService := service.NewUserService(s)

	userHandler := handlers.NewUserHandler(userService)

	r.GET("/", IndexRoute)
	r.GET("/health", HealthRoute)

	users := r.Group("/api/users")

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
