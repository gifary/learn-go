package state

import (
	"github.com/gin-gonic/gin"
	"learngo/models"
)

type Usecase interface {
	Fetch(c *gin.Context)(res []*models.State,err error)
}
