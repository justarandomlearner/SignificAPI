package main

import (
	"github.com/gin-gonic/gin"
	"github.com/justarandomlearner/SignificAPI/cmd/api/handlers"
)

func main() {

	g := gin.Default()

	g.GET("/sayhello", handlers.SayHello)

	g.GET("/api/:word", handlers.WordHandler)

	g.Run(":3000")
}
