package entity1

import (
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Username  string    `json: "user_name"`
	Email     string    `json: "email"`
	Password  string    `json: "password"`
	Age       int       `json: "age"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Username    string `json:"username"`
	TokenString string `json:"token"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}
