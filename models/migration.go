package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Model
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name"`
	Username string `json:"username" gorm:"type:varchar(128);unique_index" binding:"required"`  //set username to unique and not null
	Password string `json:"password" binding:"required"`
	Email string `json:"email" gorm:"type:varchar(128);unique_index" binding:"required"`
}

type UserTemp struct {
	ID        uint
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"` //set username to unique and not null
	Email     string `json:"email"`
}
