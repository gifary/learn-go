package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Base struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type BaseWithoutUuid struct {
	ID        uuid.UUID `gorm:"type:int;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuidString := uuid.NewV4()
	return scope.SetColumn("ID", uuidString)
}

type User struct {
	Base
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name"`
	Username string `json:"username" gorm:"type:varchar(128);unique_index" binding:"required"`  //set username to unique and not null
	Password string `json:"password" binding:"required"`
	Email string `json:"email" gorm:"type:varchar(128);unique_index" binding:"required"`
	RoleId uuid.UUID `json:"role_id" gorm:"type:char(36);not null;foreignkey:id;association_foreignkey:role_id"`
}

type Role struct {
	Base
	Name string `json:"title" binding:"required" gorm:"type:varchar(128)"`
}

type Gender string

type Customer struct {
	Base
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name"`
	Password string `json:"password" binding:"required"`
	Email string `json:"email" gorm:"type:varchar(128);"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(20);unique_index" binding:"required"`
	GoogleId string `json:"google_id" gorm:"type:varchar(128);"`
	FacebookId string `json:"facebook_id" gorm:"type:varchar(128);"`
	Gender Gender `json:"gender" gorm:"type:ENUM('male','female'); default:'female'"`
	IsActive bool `json:"is_active" gorm:"type:boolean; default:true;not null"`
}

type Midwife struct {
	Base
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name"`
	Password string `json:"password" binding:"required"`
	Email string `json:"email" gorm:"type:varchar(128);"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(20);unique_index" binding:"required"`
	KtpNumber string `json:"ktp_number" gorm:"type:varchar(20);unique_index" binding:"required"`
	GoogleId string `json:"google_id" gorm:"type:varchar(128);"`
	FacebookId string `json:"facebook_id" gorm:"type:varchar(128);"`
	Gender Gender `json:"gender" gorm:"type:ENUM('male','female'); default:'female'"`
	IsActive bool `json:"is_active" gorm:"type:boolean; default:false;not null"`
	StateId uint `json:"state_id" gorm:"type:int;not null;"`
	CityId uint `json:"city_id" gorm:"type:int;column:not null;"`
}

type State struct {
	BaseWithoutUuid
	Name string `json:"name"`
}

type City struct {
	BaseWithoutUuid
	Name string `json:"name"`
	StateId uint `json:"state_id" gorm:"type:int;not null;"`
}

type Status string

type ServiceCategory struct {
	Base
	Name string `json:"name" binding:"required" gorm:"type:varchar(128);not null;"`
	Icon string `json:"icon" gorm:"type:varchar(64);"`
	Banner string `json:"banner" gorm:"type:varchar(64);"`
}

type Service struct {
	Base
	Name string `json:"name" binding:"required" gorm:"type:varchar(128);not null;"`
	Slug string `json:"slug" gorm:"type:varchar(256);not null;"`
	Description string `json:"description" binding:"required" gorm:"type:text;"`
	Icon string `json:"icon" gorm:"type:varchar(64);"`
	Banner string `json:"banner" gorm:"type:varchar(64);"`
	Duration uint `json:"banner" binding:"required" gorm:"type:int;"`
	Note string `json:"note" gorm:"type:text;"`
	Price uint `json:"price" binding:"required" gorm:"type:int;"`
	Status Status `json:"status" gorm:"type:ENUM('active','inactive'); default:'active'"`
	ServiceCategoryId uuid.UUID `json:"service_category_id" gorm:"type:char(36);not null;foreignkey:id;association_foreignkey:service_category_id"`
}