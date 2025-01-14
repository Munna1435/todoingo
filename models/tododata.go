package models

type TodoData struct {
	Id          string `json:"id"`
	Task        string `json:"task"`
	IsCompleted bool   `json:"isCompleted"`
	UserId      string `json:"userId"`
	IsDeleted   bool   `json:"isDeleted"`
}
