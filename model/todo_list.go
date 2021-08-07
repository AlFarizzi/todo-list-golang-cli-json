package model

type Todo struct {
	Id         int    `json:"id"`
	Author     string `validate:"required" json:"author"`
	Todo       string `json:"todo"`
	Done       string `json:"done"`
	Created_At string `json:"created_at"`
}
