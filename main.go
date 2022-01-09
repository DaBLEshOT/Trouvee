package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	connectDatabase()
	token := os.Getenv("API_TOKEN")

	r.PUT("/found", tokenAuth(token), confirmCoords)
	r.POST("/add", tokenAuth(token), createTreasure)
	r.GET("/closest", tokenAuth(token), closestTreasure)

	r.Run()
}
