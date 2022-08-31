package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a api) allReleases(c *gin.Context) {
	log.Println("GET /orders")
	log.Printf("ClientIP: %s\n", c.ClientIP())

	releases, err := a.ListReleasesService.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, releases)
}
