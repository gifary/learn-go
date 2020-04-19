package usecase

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"learngo/models"
	"learngo/user"
	"time"
)

type userUsecase struct {
	userRepo user.Repository
	contextTimeout time.Duration
}

func NewUserUsecase(user user.Repository, timeout time.Duration) user.Usecase {
	return &userUsecase{
		userRepo:user,
		contextTimeout:timeout,
	}
}

func (u userUsecase) Fetch(c *gin.Context) (res []*models.User, err error) {
	return u.userRepo.Fetch(c)
}

func (u userUsecase) GetById(c *gin.Context, id int) (user *models.User, err error) {
	return u.userRepo.GetById(c,id)
}

func (u userUsecase) GetByUsername(c *gin.Context, username string) (user *models.User, err error) {
	return u.userRepo.GetByUsername(c,username)
}

func (u userUsecase) GetByEmail(c *gin.Context, email string) (user *models.User, err error) {
	return u.userRepo.GetByUsername(c,email)
}

func (u userUsecase) Update(c *gin.Context, user *models.User) error {
	return u.userRepo.Update(c,user)
}

func (u userUsecase) Store(c *gin.Context, user *models.User) error {
	return u.userRepo.Store(c,user)
}

func (u userUsecase) Delete(c *gin.Context, id uuid.UUID) error {
	return u.userRepo.Delete(c,id)
}

