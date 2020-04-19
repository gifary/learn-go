package repository

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	Helper "learngo/helper"
	"learngo/models"
	"learngo/user"
	"time"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(Conn *gorm.DB) user.Repository  {
	return &mysqlUserRepository{Conn:Conn}
}

func (m mysqlUserRepository) Fetch(c *gin.Context) (res []*models.User, err error) {
	var users []*models.User
	var errorDb = m.Conn.Find(&users).Error

	return users, errorDb
}

func (m mysqlUserRepository) GetById(c *gin.Context, id int) (user *models.User, err error) {
	var u *models.User
	var errorDb = m.Conn.Where("id=?",id).First(&u).Error

	return u,errorDb
}

func (m mysqlUserRepository) GetByUsername(c *gin.Context, username string) (user *models.User, err error) {
	var u *models.User
	var errorDb = m.Conn.Where("username=?",username).First(&u).Error

	return u,errorDb
}

func (m mysqlUserRepository) GetByEmail(c *gin.Context, email string) (user *models.User, err error) {
	var u *models.User
	var errorDb = m.Conn.Where("email=?",email).First(&u).Error

	return u,errorDb
}

func (m mysqlUserRepository) Update(c *gin.Context, user *models.User) error {
	var u models.User

	var errDb = m.Conn.Where("id=?",user.ID).First(&u).Error

	var checkUser models.User
	m.Conn.Where("email=?",user.Email).First(&checkUser)

	if checkUser.ID != u.ID{
		return errors.New("email already exist")
	}

	m.Conn.Where("username=?",user.Username).First(&checkUser)

	if checkUser.ID != u.ID {
		return errors.New("username already exist")
	}

	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Username = user.Username
	u.Email = user.Email
	if len(user.Password)>0{
		u.Password = Helper.HashPassword(user.Password)
	}else{
		u.Password = u.Password
	}

	if errDb!=nil {
		return errDb
	}else{
		m.Conn.Save(&u)
	}

	return nil

}

func (m mysqlUserRepository) Store(c *gin.Context, user *models.User) error {
	//cek username and email first

	var checkUser models.User
	m.Conn.Where("email=?",user.Email).First(&checkUser)

	if checkUser.Email!=""{
		return errors.New("email already exist")
	}

	m.Conn.Where("username=?",user.Username).First(&checkUser)

	if checkUser.Username!=""{
		return errors.New("username already exist")
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Password = Helper.HashPassword(user.Password)

	m.Conn.Save(&user)

	return nil
}

func (m mysqlUserRepository) Delete(c *gin.Context, id uuid.UUID) error {
	var u models.User
	var errDb = m.Conn.Where("id=?",id).Delete(&u).Error
	if errDb != nil {
		return errDb
	}else{
		return nil
	}
}

