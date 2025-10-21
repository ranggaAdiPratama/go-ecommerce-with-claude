package server

import (
	"database/sql"
	"log"
	"ranggaAdiPratama/go-with-claude/internal/config"
	"ranggaAdiPratama/go-with-claude/internal/database"
	"ranggaAdiPratama/go-with-claude/internal/routes"
	"ranggaAdiPratama/go-with-claude/internal/utils"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config            *config.Config
	router            *gin.Engine
	store             *database.Store
	pasetoMaker       *utils.PasetoMaker
	cloudinaryService *utils.CloudinaryService
}

func New(db *sql.DB, cfg *config.Config) *Server {
	gin.SetMode(cfg.GinMode)

	store := database.NewStore(db)
	router := gin.Default()

	pasetoMaker, err := utils.NewPasetoMaker(cfg.TokenSymmetricKey)

	if err != nil {
		log.Fatal("Cannot create PASETO maker:", err)
	}

	cloudinaryService, err := utils.NewCloudinaryService(cfg.CloudinaryName, cfg.CloudinaryAPIKey, cfg.CloudinaryAPISecret, cfg.CloudinaryFolder)

	if err != nil {
		log.Fatal("Cannot set Cloudinary Service:", err)
	}

	s := &Server{
		config:            cfg,
		router:            router,
		store:             store,
		pasetoMaker:       pasetoMaker,
		cloudinaryService: cloudinaryService,
	}

	routes.Index(router, store, pasetoMaker, cfg, cloudinaryService)

	return s
}

func (s *Server) Start() error {
	return s.router.Run(":" + s.config.Port)
}
