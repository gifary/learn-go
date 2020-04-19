package usecase

import (
	"github.com/gin-gonic/gin"
	"learngo/city"
	"learngo/models"
)

type cityUsecase struct {
	cityRepository city.Repository
}

func (c2 cityUsecase) Fetch(c *gin.Context) (res []*models.City, err error) {
	return c2.cityRepository.Fetch(c)
}

func (c2 cityUsecase) GetByStateId(c *gin.Context, id int) (res []*models.City, err error) {
	return c2.cityRepository.GetByStateId(c,id)
}

func NewStateUsecase(city city.Repository) city.Usecase {
	return &cityUsecase{cityRepository:city}
}

