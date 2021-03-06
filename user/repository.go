package user

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"learngo/models"
)

type Repository interface {
	Fetch(c *gin.Context) (res []*models.User,err error)
	GetById(c *gin.Context,id int) (user *models.User, err error)
	GetByUsername(c *gin.Context, username string)(user *models.User, err error)
	GetByEmail(c *gin.Context,email string)(user *models.User,err error)
	Update(c *gin.Context,user *models.User) error
	Store(c *gin.Context, user *models.User) error
	Delete(c *gin.Context,id uuid.UUID) error
}
