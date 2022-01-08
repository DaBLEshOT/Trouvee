package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	createDatabase()

	r := gin.Default()
	r.POST("/found", confirmCoords)
	r.POST("/add", createTreasure)
	r.Run()
}
