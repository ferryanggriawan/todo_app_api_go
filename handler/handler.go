package handler

import (
	"todo_app_api_go/model"

	"github.com/gin-gonic/gin"
)

func RootHandler(ctx *gin.Context) {
	resp := model.Default{
		Message: "This is root path",
	}

	ctx.JSON(200, resp)
}
