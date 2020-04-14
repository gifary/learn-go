package http

import (
	"github.com/gin-gonic/gin"
	"learngo/course"
	"learngo/helper"
	"learngo/models"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseArraySuccess struct {
	Message string `json:"message"`
	Data []*models.Course `json:"data"`
}

type ResponseObjectSuccess struct {
	Message string `json:"message"`
	Data *models.Course `json:"data"`
}

type CourseHandler struct {
	CourseUsecase course.Usecase
}

func NewCourseHandler(r *gin.Engine, cu course.Usecase)  {
	handler := &CourseHandler{CourseUsecase:cu}
	r.GET("course-by-user",handler.GetByUser)
}

func (ch *CourseHandler) GetByUser(c *gin.Context)  {
	listArr,err := ch.CourseUsecase.GetByUser(c)

	if err != nil {
		c.JSON(helper.GetStatusCode(err),ResponseError{Message:err.Error()})
		return
	}

	c.JSON(http.StatusOK, ResponseArraySuccess{Message:"Success get",Data:listArr})
}