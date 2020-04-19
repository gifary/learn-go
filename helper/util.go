package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"learngo/models"
	"net/http"
	"os"
	"strconv"
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

func GetCurrentUserId(c* gin.Context) uuid.UUID  {
	// claims := jwt.ExtractClaims(c)
	user,_ := c.Get(os.Getenv("IDENTIFIED_KEY"))

	return user.(*models.User).ID
}

func StringToUInt(s string) uint {
	value,err := strconv.ParseUint(s,10,64)
	if  err != nil {
		fmt.Print(err)
	}

	return uint(value)
}

func StringToInt(s string) int {
	value,err := strconv.ParseInt(s,10,64)
	if  err != nil {
		fmt.Print(err)
	}

	return int(value)
}
