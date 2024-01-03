package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	rword "github.com/justarandomlearner/SignificAPI/internal/repository"
	sword "github.com/justarandomlearner/SignificAPI/internal/service"
)

func SayHello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func WordHandler(ctx *gin.Context) {
	service := &sword.Service{
		Repository: rword.Repository{
			// Conn: conn,
		},
	}
	word := ctx.Param("word")

	if strings.TrimSpace(word) == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("empty word"),
		})
		return
	}

	outerContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	payload, err := service.FindWord(outerContext, word)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, payload)
}
