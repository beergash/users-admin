package model

import (
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	BirthDate time.Time `json:"birthdate"`
}

type UserSearchPaginator struct {
	TotalElements int    `json:"totalelements"`
	Page          int    `json:"page"`
	Users         []User `json:"users"`
}

type UserSearchRequest struct {
	Page    int    `json:"page"`
	Size    int    `json:"size"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
