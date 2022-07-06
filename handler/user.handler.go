package handler

import (
	"todo_app_api_go/model"

	"github.com/gin-gonic/gin"
)

func UserHandler(ctx *gin.Context) {
	resp := model.Default{
		Message: "This is user path",
	}

	ctx.JSON(200, resp)
}
