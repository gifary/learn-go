package usecase

import (
	"github.com/gin-gonic/gin"
	"learngo/models"
	"learngo/state"
)

type stateUsecase struct {
	stateRepo state.Repository
}

func (s stateUsecase) Fetch(c *gin.Context) (res []*models.State, err error) {
	return s.stateRepo.Fetch(c)
}

func NewStateUsecase(state state.Repository) state.Usecase {
	return &stateUsecase{stateRepo:state}
}