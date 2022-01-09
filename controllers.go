package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	maxDistance = 10
)

func confirmCoords(c *gin.Context) {
	var point Point
	var treasures []Treasure

	if err := c.ShouldBind(&point); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	DB.Find(&treasures)
	var treasure Treasure
	for _, t := range treasures {
		p := NewPoint(t.Lat, t.Lng)
		if distance := point.GreatCircleDistance(p); distance < maxDistance {
			DB.Model(&t).Update("Found", true)
			treasure = t
			break
		}
	}

	c.JSON(http.StatusOK, treasure)
}

func createTreasure(c *gin.Context) {
	var treasure Treasure

	if err := c.ShouldBind(&treasure); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	DB.Create(&treasure)

	c.JSON(http.StatusOK, treasure)
}
