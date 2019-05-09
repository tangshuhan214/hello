package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
	Sex  string `json:"sex"`
}

