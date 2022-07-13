package auth_service

import (
	"encoding/json"
	"todo_app_api_go/model"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) *model.User {
	defer handleError()

	userForm := checkRequest(c)
	user := insertToDb(userForm)

	return user
}

func checkRequest(c *gin.Context) *model.UserForm {
	user := model.UserForm{}

	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(&user); err != nil {
		panic("Bad Request")
	}

	return &user
}

func insertToDb(m *model.UserForm) *model.User {
	user := &model.User{}

	return user
}

func handleError() {

}
