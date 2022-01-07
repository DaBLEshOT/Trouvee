package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var points = []Point{
	{
		Lat: 46.561916,
		Lng: 15.638865,
	},
}

func main() {
	r := gin.Default()
	r.POST("/found", confirmCoords)
	r.Run()
}

func confirmCoords(c *gin.Context) {
	var point Point

	if err := c.BindJSON(&point); err != nil {
		log.Fatalln(err)
		return
	}

	found := false
	for _, p := range points {
		if distance := point.GreatCircleDistance(&p); distance < 10 {
			found = true
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"found": found,
	})
}
