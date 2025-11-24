package routes

import (
	"net/http"
	"ranggaAdiPratama/go-with-claude/internal/config"
	"ranggaAdiPratama/go-with-claude/internal/database"
	"ranggaAdiPratama/go-with-claude/internal/handlers"
	"ranggaAdiPratama/go-with-claude/internal/responses"
	"ranggaAdiPratama/go-with-claude/internal/service"
	"ranggaAdiPratama/go-with-claude/internal/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Index(r *gin.Engine, s *database.Store, p *utils.PasetoMaker, c *config.Config, cl *utils.CloudinaryService) {
	authService := service.NewAuthService(s, p, c)
	categoryService := service.NewCategoryService(s)
	settingService := service.NewSettingService(s, cl)
	shopService := service.NewShopService(s, cl)
	userService := service.NewUserService(s)

	authHandler := handlers.NewAuthHandler(authService, userService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	settingHandler := handlers.NewSettingHandler(settingService)
	shopHandler := handlers.NewShopHandler(shopService)
	userHandler := handlers.NewUserHandler(userService)

	r.Use(cors.Default())

	r.GET("/", IndexRoute)
	r.GET("/health", HealthRoute)

	authRoute(r, authHandler, p)
	categoryRoute(r, categoryHandler, p)
	settingRoute(r, settingHandler, p)
	shopRoute(r, shopHandler, p)
	userRoute(r, userHandler, p)
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
