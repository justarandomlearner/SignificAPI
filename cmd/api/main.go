package main

import (
	"github.com/gin-gonic/gin"
	"github.com/justarandomlearner/SignificAPI/internal/handlers"
)

func main() {

	g := gin.Default()

	g.GET("/api/:word", handlers.WordHandler)

	g.Run(":3000")
}
