package handler

import (
	"github.com/gin-gonic/gin"
)

func RootHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "root path",
	})
}
