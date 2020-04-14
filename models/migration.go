package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name"`
	Username string `json:"username" gorm:"type:varchar(128);unique_index" binding:"required"`  //set username to unique and not null
	Password string `json:"password" binding:"required"`
	Email string `json:"email" gorm:"type:varchar(128);unique_index" binding:"required"`
	Courses []Course `json:"courses" gorm:"foreignkey:UserRefer"`
}

type UserTemp struct {
	ID        uint
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"` //set username to unique and not null
	Email     string `json:"email"`
}

type Course struct {
	Model
	Title string `json:"title" binding:"required" gorm:"type:varchar(128)"`
	Description string `json:"description" gorm:"type:text"`
	User      User `json:"-" gorm:"foreignkey:UserRefer"` // use UserRefer as foreign key
	UserRefer uint `json:"-"`
}