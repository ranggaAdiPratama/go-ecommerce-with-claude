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

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Index(r *gin.Engine, s *database.Store, p *utils.PasetoMaker, c *config.Config, cl *utils.CloudinaryService) {
	authService := service.NewAuthService(s, p, c)
	settingService := service.NewSettingService(s, cl)
	shopService := service.NewShopService(s, cl)
	userService := service.NewUserService(s)

	authHandler := handlers.NewAuthHandler(authService, userService)
	settingHandler := handlers.NewSettingHandler(settingService)
	shopHandler := handlers.NewShopHandler(shopService)
	userHandler := handlers.NewUserHandler(userService)

	r.Use(cors.Default())

	r.GET("/", IndexRoute)
	r.GET("/health", HealthRoute)

	auth := r.Group("/api/auth")
	myShop := r.Group("/api/my-shop").Use(middleware.AuthMiddleware(p))
	setting := r.Group("/api/settings")
	shops := r.Group("/api/shops")
	users := r.Group("/api/users").Use(middleware.AuthMiddleware(p)).Use(middleware.RequireRole("admin"))

	auth.POST("/login", authHandler.Login)
	auth.POST("/logout", middleware.AuthMiddleware(p), authHandler.Logout)
	auth.POST("/refresh-token", authHandler.RefreshToken)
	auth.POST("/register", authHandler.Register)

	myShop.POST("", shopHandler.Store)
	myShop.PUT("", middleware.RequireRole("user"), shopHandler.UpdatePersonal)

	setting.GET("", settingHandler.Index)
	setting.POST("", middleware.AuthMiddleware(p), middleware.RequireRole("admin"), settingHandler.StoreOrUpdate)

	shops.GET("", shopHandler.Index)

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
