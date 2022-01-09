package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	maxDistance = 6
)

func confirmCoords(c *gin.Context) {
	var point Point
	var treasures []Treasure

	if err := c.ShouldBind(&point); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	DB.Where("found = FALSE").Find(&treasures)
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

func closestTreasure(c *gin.Context) {
	var point Point
	var treasures []Treasure

	if err := c.ShouldBind(&point); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	DB.Where("found = FALSE").Find(&treasures)
	var closestDistance float64
	var closestTreasure Treasure
	for _, t := range treasures {
		p := NewPoint(t.Lat, t.Lng)
		if distance := point.GreatCircleDistance(p); distance < closestDistance {
			closestDistance = distance
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"distance": closestDistance,
		"hint":     closestTreasure.Hint,
	})
}

func tokenAuth(apiToken string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		if bearer == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token required"})
			return
		}

		token := strings.Replace(bearer, "Bearer ", "", 1)
		if token != apiToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.Next()
	}
}
