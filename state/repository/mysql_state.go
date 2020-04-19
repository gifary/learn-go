package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"learngo/models"
	"learngo/state"
)

type mysqlStateRepository struct {
	Conn *gorm.DB
}

func (m mysqlStateRepository) Fetch(c *gin.Context) (res []*models.State, err error) {
	var states []*models.State
	var errorDb = m.Conn.Find(&states).Error

	return states,errorDb
}

func NewMysqlStateRepository(Conn *gorm.DB) state.Repository  {
	return &mysqlStateRepository{Conn:Conn}
}