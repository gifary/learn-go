package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"learngo/city"
	"learngo/models"
)

type mysqlCityRepository struct {
	Conn *gorm.DB
}

func (m mysqlCityRepository) GetByStateId(c *gin.Context, id int) (res []*models.City, err error) {
	var u []*models.City
	var errDb = m.Conn.Where("state_id=?",id).Find(&u).Error

	return u,errDb

}

func (m mysqlCityRepository) Fetch(c *gin.Context) (res []*models.City, err error) {
	var states []*models.City
	var errorDb = m.Conn.Find(&states).Error

	return states,errorDb
}

func NewMysqlStateRepository(Conn *gorm.DB) city.Repository  {
	return &mysqlCityRepository{Conn:Conn}
}