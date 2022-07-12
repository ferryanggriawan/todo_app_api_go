package root_handler

import (
	"todo_app_api_go/handler"
	"todo_app_api_go/model"

	"github.com/gin-gonic/gin"
)

func RootHandler(ctx *gin.Context) {
	defer handler.ErrorHttp(ctx)

	resp := model.Default{
		Message: "This is root path",
	}

	ctx.JSON(200, resp)
}
