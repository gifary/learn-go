package city

import (
	"github.com/gin-gonic/gin"
	"learngo/models"
)

type Repository interface {
	Fetch(c *gin.Context)(res []*models.City,err error)
	GetByStateId(c *gin.Context, id int)(res []*models.City,err error)
}
