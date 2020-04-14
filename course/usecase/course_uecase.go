package usecase

import (
	"github.com/gin-gonic/gin"
	"learngo/course"
	"learngo/models"
)

type course_usecase struct {
	courseRepo course.Repistory
}

func (c2 course_usecase) GetByUser(c *gin.Context) (res []*models.Course, err error) {
	return c2.courseRepo.GetByUser(c)
}

func (c2 course_usecase) Store(c *gin.Context, course models.Course) error {
	panic("implement me")
}

func (c2 course_usecase) Update(c *gin.Context, course *models.Course) error {
	panic("implement me")
}

func (c2 course_usecase) Delete(c *gin.Context, id int) error {
	panic("implement me")
}

func NewCourseUsecase(course course.Repistory) course.Usecase  {
	return &course_usecase{courseRepo:course}
}
