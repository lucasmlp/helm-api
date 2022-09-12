package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/machado-br/helm-api/adapters/models"
)

func (a api) allReleases(c *gin.Context) {
	log.Println("GET /")

	releases, err := a.listReleasesService.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, releases)
}

func (a api) installChart(c *gin.Context) {
	log.Println("POST /?dryRun")

	var chartPayload InstallChartPayload
	err := c.ShouldBindJSON(&chartPayload)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	dryRun, err := strconv.ParseBool(c.Query("dryRun"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	chart := models.Chart{
		Name:    chartPayload.ChartName,
		Version: chartPayload.Version,
		RepoURL: chartPayload.RepoURL,
	}

	err = a.installChartService.Run(chartPayload.ReleaseName, dryRun, chart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusCreated)
	return
}

func (a api) uninstallRelease(c *gin.Context) {
	log.Println("DELETE /name?dryRun")

	releaseName := c.Param("name")
	dryRun, err := strconv.ParseBool(c.Query("dryRun"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = a.uninstallReleaseService.Run(releaseName, dryRun)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
	return
}
