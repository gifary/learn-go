package course

import (
	"github.com/gin-gonic/gin"
	"learngo/models"
)

type Repistory interface {
	GetByUser(c *gin.Context) (res []*models.Course,err error)
	Store(c *gin.Context,course models.Course) error
	Update(c *gin.Context,course *models.Course) error
	Delete(c *gin.Context,id int) error
}
