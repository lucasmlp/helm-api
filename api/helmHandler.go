package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/machado-br/helm-api/adapters/models"
)

func (a api) allReleases(c *gin.Context) {
	log.Println("GET /helm")

	releases, err := a.listReleasesService.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, releases)
}

func (a api) installChart(c *gin.Context) {
	log.Println("POST /helm")

	var chartPayload InstallChartPayload
	err := c.ShouldBindJSON(&chartPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	
	chart := models.Chart{
		Name: chartPayload.ChartName,
		Version: chartPayload.Version,
		RepoURL: chartPayload.RepoURL,
	}

	err = a.installChartService.Run(chartPayload.ReleaseName, chartPayload.DryRun, chart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "")
}
