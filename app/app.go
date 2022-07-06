package app

import (
	"fmt"
	_ "todo_app_api_go/app/database"

	"github.com/gin-gonic/gin"
)

func Print(abc interface{}) {
	fmt.Printf("%s", abc)
}

func ErrorHttp(ctx *gin.Context) {
	if ctx.Errors != nil {
		resp := gin.H{
			"message": "unknown error",
		}

		ctx.JSON(500, resp)
	}
}
