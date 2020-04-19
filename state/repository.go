package state

import (
	"github.com/gin-gonic/gin"
	"learngo/models"
)

type Repository interface {
	Fetch(c *gin.Context)(res []*models.State,err error)
}
