package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	maxDistance = 10
)

func confirmCoords(c *gin.Context) {
	var point Point
	var treasures []Treasure
	db := openDb()

	if err := c.Bind(&point); err != nil {
		log.Fatalln(err)
		return
	}

	db.Find(&treasures)
	var treasure Treasure
	for _, t := range treasures {
		p := NewPoint(t.Lat, t.Lng)
		if distance := point.GreatCircleDistance(p); distance < maxDistance {
			db.Model(&t).Update("Found", true)
			treasure = t
			break
		}
	}

	c.JSON(http.StatusOK, treasure)
}

func createTreasure(c *gin.Context) {
	var treasure Treasure
	db := openDb()

	if err := c.Bind(&treasure); err != nil {
		log.Fatalln(err)
		return
	}

	db.Create(&treasure)

	c.JSON(http.StatusOK, treasure)
}