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
	config      *config.Config
	router      *gin.Engine
	store       *database.Store
	pasetoMaker *utils.PasetoMaker
}

func New(db *sql.DB, cfg *config.Config) *Server {
	gin.SetMode(cfg.GinMode)

	store := database.NewStore(db)
	router := gin.Default()

	pasetoMaker, err := utils.NewPasetoMaker(cfg.TokenSymmetricKey)

	if err != nil {
		log.Fatal("Cannot create PASETO maker:", err)
	}

	s := &Server{
		config:      cfg,
		router:      router,
		store:       store,
		pasetoMaker: pasetoMaker,
	}

	routes.Index(router, store, pasetoMaker, cfg)

	return s
}

func (s *Server) Start() error {
	return s.router.Run(":" + s.config.Port)
}
