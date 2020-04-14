package helper

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"learngo/models"
	"net/http"
	"os"
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func GetCurrentUserId(c* gin.Context) uint  {
	// claims := jwt.ExtractClaims(c)
	user,_ := c.Get(os.Getenv("IDENTIFIED_KEY"))

	return user.(*models.User).ID
}