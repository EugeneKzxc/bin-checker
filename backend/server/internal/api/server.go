package api

import (
	"fmt"
	"log"

	"backend/server/internal/api/gateway"
	"backend/server/internal/config"
	"backend/server/internal/repo"
	"backend/server/internal/scope"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router interface {
	NewRouter(engine *gin.Engine, cfg *config.Config, storage *scope.Storage, repository *repo.Repository)
}

type Server struct {
	engine  *gin.Engine
	cfg     *config.Config
	storage *scope.Storage
	repo    *repo.Repository
}

func InitServer(engine *gin.Engine, cfg *config.Config, storage *scope.Storage, repo *repo.Repository) *Server {
	server := &Server{
		engine:  engine,
		cfg:     cfg,
		storage: storage,
		repo:    repo,
	}

	server.initRouter()
	return server
}

func NewRouters(routers []Router, engine *gin.Engine, cfg *config.Config, storage *scope.Storage, repository *repo.Repository) {
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	}))
	for _, router := range routers {
		router.NewRouter(engine, cfg, storage, repository)
	}
}

func (server *Server) initRouter() {
	NewRouters([]Router{
		&gateway.Router{},
	}, server.engine, server.cfg, server.storage, server.repo)
}

func (server *Server) Run(port uint) {
	log.Println("Gin server started")
	if err := server.engine.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Println("server run error : ", err)
	}
}
