package main

import (
	"log"

	"backend/server/internal/api"
	"backend/server/internal/config"
	"backend/server/internal/repo"
	"backend/server/internal/scope"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := new(config.Config)
	if err := cfg.Read(); err != nil {
		log.Fatal(err)
	}

	r, err := repo.NewRepository(cfg)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	data, err := r.GetAllBins()
	if err != nil {
		log.Fatal("Error on reading data from db: ", err)
	}
	storage := scope.NewStorage(data)

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Logger())

	api.InitServer(engine, cfg, storage, r).Run(cfg.Port)
}
