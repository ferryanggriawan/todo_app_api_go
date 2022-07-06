package handler

import (
	"fmt"
	"todo_app_api_go/app"
	"todo_app_api_go/app/service/auth"
	"todo_app_api_go/model"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(ctx *gin.Context) {
	defer app.ErrorHttp(ctx)

	resp := model.Default{
		Message: "This is auth path",
	}

	test := auth.RegisterUser()

	fmt.Println(test)

	ctx.JSON(200, resp)
}
