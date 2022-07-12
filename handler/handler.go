package handler

import "github.com/gin-gonic/gin"

func ErrorHttp(ctx *gin.Context) {
	if ctx.Errors != nil {
		resp := gin.H{
			"message": "unknown error",
		}

		ctx.JSON(500, resp)
	}
}
