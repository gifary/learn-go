package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"learngo/course"
	"learngo/helper"
	"learngo/models"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func (m mysqlUserRepository) GetByUser(c *gin.Context) (res []*models.Course, err error) {
	var courses []*models.Course
	var userId =helper.GetCurrentUserId(c)

	var errorDb = m.Conn.Where("user_refer=?",userId).Find(&courses).Error

	return courses, errorDb
}

func (m mysqlUserRepository) Store(c *gin.Context, course models.Course) error {
	panic("implement me")
}

func (m mysqlUserRepository) Update(c *gin.Context, course *models.Course) error {
	panic("implement me")
}

func (m mysqlUserRepository) Delete(c *gin.Context, id int) error {
	panic("implement me")
}

func NewMysqlCourseRepository(Conn *gorm.DB) course.Repistory  {
	return &mysqlUserRepository{Conn:Conn}
}