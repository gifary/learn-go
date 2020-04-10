package http

import (
	"github.com/gin-gonic/gin"
	"learngo/helper"
	"learngo/models"
	"learngo/user"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseArraySuccess struct {
	Message string `json:"message"`
	Data []*models.User `json:"data"`
}

type ResponseObjectSuccess struct {
	Message string `json:"message"`
	Data *models.User `json:"data"`
}

type UserHandler struct {
	UserUsecase user.Usecase
}

func NewUserHandler(r *gin.Engine, us user.Usecase)  {
	handler := &UserHandler{UserUsecase:us}

	r.GET("/users",handler.FetchUser)
	r.GET("/users/:id/by-id",handler.GetById)
	r.POST("/users",handler.Store)
	r.PUT("/users/:id",handler.Update)
}

func (u *UserHandler) FetchUser(c *gin.Context) {

	listArr,err := u.UserUsecase.Fetch(c)

	if err != nil {
		c.JSON(helper.GetStatusCode(err),ResponseError{Message:err.Error()})
		return
	}

	c.JSON(http.StatusOK, ResponseArraySuccess{Message:"Success get",Data:listArr})
}

func (u *UserHandler) GetById(c *gin.Context)  {
	var id, _ = strconv.Atoi(c.Params.ByName("id"))

	user,err := u.UserUsecase.GetById(c,id)

	if err !=nil{
		c.JSON(helper.GetStatusCode(err),ResponseError{Message:err.Error()})
		return
	}

	c.JSON(http.StatusOK, ResponseObjectSuccess{Data:user,Message:"success"})
}

func (u *UserHandler) GetByUsername(c *gin.Context)  {
	var username =  c.Params.ByName("username")
	user,err := u.UserUsecase.GetByUsername(c,username)

	if err !=nil{
		c.JSON(helper.GetStatusCode(err),ResponseError{Message:err.Error()})
		return
	}

	c.JSON(http.StatusOK, ResponseObjectSuccess{Data:user,Message:"success"})
}

func (u *UserHandler) GetByEmail(c *gin.Context)  {
	var email =  c.Params.ByName("email")

	user,err := u.UserUsecase.GetByEmail(c,email)

	if err !=nil{
		c.JSON(helper.GetStatusCode(err),ResponseError{Message:err.Error()})
		return
	}

	c.JSON(http.StatusOK, ResponseObjectSuccess{Data:user,Message:"success"})
}

func (u *UserHandler)Store(c* gin.Context)  {
	var user models.User

	if err := c.ShouldBind(&user); err!=nil{
		c.JSON(helper.GetStatusCode(err),ResponseError{Message:err.Error()})
		return
	}

	err := u.UserUsecase.Store(c,&user)

	if err != nil{
		c.JSON(helper.GetStatusCode(err),ResponseError{Message:err.Error()})
		return
	}

	c.JSON(http.StatusOK, ResponseObjectSuccess{Data:&user,Message:"success"})
	return
}

func(u *UserHandler) Update(c *gin.Context)  {
	var user models.User
	var id, _ = strconv.ParseUint(c.Params.ByName("id"),10,64)

	if err := c.ShouldBind(&user); err!=nil{
		c.JSON(helper.GetStatusCode(err),ResponseError{Message:err.Error()})
		return
	}

	user.ID = uint(id)

	err := u.UserUsecase.Update(c,&user)

	if err != nil{
		c.JSON(helper.GetStatusCode(err),ResponseError{Message:err.Error()})
		return
	}

	c.JSON(http.StatusOK, ResponseObjectSuccess{Data:&user,Message:"success"})
	return
}