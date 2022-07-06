package model

type Task struct {
	Id        int    `json:"id"`
	Task      string `json:"task"`
	Task_Id   int    `json:"task_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
