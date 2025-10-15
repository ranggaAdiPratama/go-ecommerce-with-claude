package server

import (
	"database/sql"
	"ranggaAdiPratama/go-with-claude/internal/config"
	"ranggaAdiPratama/go-with-claude/internal/database"
	"ranggaAdiPratama/go-with-claude/internal/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config *config.Config
	router *gin.Engine
	store  *database.Store
}

func New(db *sql.DB, cfg *config.Config) *Server {
	gin.SetMode(cfg.GinMode)

	store := database.NewStore(db)
	router := gin.Default()

	s := &Server{
		config: cfg,
		router: router,
		store:  store,
	}

	routes.Index(router, store)

	return s
}

func (s *Server) Start() error {
	return s.router.Run(":" + s.config.Port)
}
