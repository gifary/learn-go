package http

import (
	"github.com/gin-gonic/gin"
	"learngo/city"
	"learngo/helper"
	"learngo/models"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseArraySuccess struct {
	Message string `json:"message"`
	Data []*models.City `json:"data"`
}

type ResponseObjectSuccess struct {
	Message string `json:"message"`
	Data *models.City `json:"data"`
}

type cityHandler struct {
	CityUsecase city.Usecase
}

func NewStateHandler(r *gin.Engine, s city.Usecase)  {
	handler := &cityHandler{CityUsecase:s}

	r.GET("/city",handler.FetchCity)
	r.GET("/city/{id}/state",handler.GetByStateId)
}

func(u *cityHandler) FetchCity(c *gin.Context)  {
	listArr,err := u.CityUsecase.Fetch(c)

	if err != nil {
		c.JSON(helper.GetStatusCode(err),ResponseError{Message:err.Error()})
		return
	}

	c.JSON(http.StatusOK, ResponseArraySuccess{Message:"Success get",Data:listArr})
}

func(u *cityHandler) GetByStateId(c *gin.Context)  {
	cityArr,err := u.CityUsecase.GetByStateId(c,helper.StringToInt(c.Params.ByName("id")))

	if err != nil {
		c.JSON(helper.GetStatusCode(err),ResponseError{Message:err.Error()})
		return
	}

	c.JSON(http.StatusOK, ResponseArraySuccess{Message:"Success get",Data:cityArr})
}