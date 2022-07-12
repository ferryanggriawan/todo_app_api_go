package auth_handler

import (
	"todo_app_api_go/handler"
	"todo_app_api_go/model"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	defer handler.ErrorHttp(ctx)

	resp := model.Default{
		Message: "This is auth path",
	}

	ctx.JSON(200, resp)
}
