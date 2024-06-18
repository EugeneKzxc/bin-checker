package gateway

import (
	"log"
	"net/http"
	"strconv"

	"backend/server/internal/config"
	"backend/server/internal/models"
	"backend/server/internal/repo"
	"backend/server/internal/scope"
	"backend/server/internal/service"

	"github.com/gin-gonic/gin"
)

type Router struct {
	cfg     *config.Config
	storage *scope.Storage
	reader  *repo.Repository
}

func (r *Router) NewRouter(engine *gin.Engine, cfg *config.Config, storage *scope.Storage, repository *repo.Repository) {
	r.cfg = cfg
	r.storage = storage
	r.reader = repository
	engine.GET("/bin", r.bin)
}

func (r *Router) bin(c *gin.Context) {
	arg := c.Query("bin")
	if arg == "" {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "missing bin parameter",
			Data:    nil,
		})
		return
	}
	bin, _ := strconv.Atoi(arg)
	bank := r.storage.Mapping[bin]
	if bank == "" && r.cfg.ExternalApi == true {
		log.Print("going to api")
		info, err := service.SearchBinInfo(bin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Status:  "error",
				Message: err.Error(),
				Data:    nil,
			})
			return
		}
		r.storage.Add(bin, info)
		err = r.reader.InsertBin(bin, info)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, models.Response{
				Status:  "error",
				Message: "internal error",
				Data:    nil,
			})
			return
		}
		bank = info
	}

	c.JSON(http.StatusOK, models.Response{
		Status: "success",
		Data:   bank,
	})
}
