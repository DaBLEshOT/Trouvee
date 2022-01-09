package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	connectDatabase()

	r.POST("/found", confirmCoords)
	r.POST("/add", createTreasure)

	r.Run()
}
