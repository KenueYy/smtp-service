package models

type Message struct {
	Code  string `json:"Code" binding:"required"`
	Email string `json:"Email" binding:"required,email"`
}
