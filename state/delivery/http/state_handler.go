package http

import (
	"github.com/gin-gonic/gin"
	"learngo/helper"
	"learngo/models"
	"learngo/state"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseArraySuccess struct {
	Message string `json:"message"`
	Data []*models.State `json:"data"`
}

type ResponseObjectSuccess struct {
	Message string `json:"message"`
	Data *models.State `json:"data"`
}

type stateHandler struct {
	StateUsecase state.Usecase
}

func NewStateHandler(r *gin.Engine, s state.Usecase)  {
	handler := &stateHandler{StateUsecase:s}

	r.GET("/states",handler.FetchState)
}

func(u *stateHandler) FetchState(c *gin.Context)  {
	listArr,err := u.StateUsecase.Fetch(c)

	if err != nil {
		c.JSON(helper.GetStatusCode(err),ResponseError{Message:err.Error()})
		return
	}

	c.JSON(http.StatusOK, ResponseArraySuccess{Message:"Success get",Data:listArr})
}