package model

type User struct {
	Id         int    `json:"id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	UserName   string `json:"user_name"`
}

type UserForm struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Username string `json:"usernames"`
	Password string `json:"password"`
}
