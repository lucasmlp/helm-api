package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/machado-br/helm-api/services/listReleases"
)

type api struct {
	ListReleasesService listReleases.Service
}

func NewApi(
	ListReleasesService listReleases.Service,
) (api, error) {
	return api{
		ListReleasesService: ListReleasesService,
	}, nil
}

func (a api) Engine() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)

	root := router.Group("")
	{
		root.GET("/ping", func(c *gin.Context) {
			log.Printf("ClientIP: %s\n", c.ClientIP())

			c.JSON(http.StatusOK, "pong")
		})
	}
	helmRoot := router.Group("/helm")
	{
		helmRoot.GET("", a.allReleases)
	}

	return router
}

func (a api) Run() {

	router := a.Engine()
	router.Run()
}
