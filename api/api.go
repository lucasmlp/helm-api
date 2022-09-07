package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/machado-br/helm-api/services/installChart"
	"github.com/machado-br/helm-api/services/listReleases"
	"github.com/machado-br/helm-api/services/uninstallRelease"
)

type api struct {
	listReleasesService listReleases.Service
	installChartService installChart.Service
	uninstallReleaseService uninstallRelease.Service
}

func NewApi(
	listReleasesService listReleases.Service,
	installChartService installChart.Service,
	uninstallReleaseService uninstallRelease.Service,
) (api, error) {
	return api{
		listReleasesService: listReleasesService,
		installChartService: installChartService,
		uninstallReleaseService: uninstallReleaseService,
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
		root.GET("/", a.allReleases)
		root.POST("/", a.installChart)
		root.DELETE("/:name", a.uninstallRelease)
	}

	return router
}

func (a api) Run() {

	router := a.Engine()
	router.Run()
}
