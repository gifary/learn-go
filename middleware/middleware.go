package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	Helper "learngo/helper"
	"learngo/models"
	"os"
	"time"
)

var user models.User
// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	Conn *gorm.DB
}

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (m *GoMiddleware)JWT() (*jwt.GinJWTMiddleware, error) {
	secretKey := os.Getenv("SECRET_KEY")

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(secretKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: os.Getenv("IDENTIFIED_KEY"),
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					os.Getenv("IDENTIFIED_KEY"): v.ID,
					"username": v.Username,
					"email": v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			userId := claims[os.Getenv("IDENTIFIED_KEY")].(float64)
			return &models.User{
				ID: uint(userId),
				Username:claims["username"].(string),
				Email:claims["email"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userName := loginVals.Username
			password := loginVals.Password

			//query to database

			var errFindUser = m.Conn.Where("username=?",userName).First(&user).Error
			matchPassword := Helper.CheckPasswordHash(password,user.Password)
			if errFindUser!=nil || !matchPassword {
				return nil, jwt.ErrFailedAuthentication
			}else{
				return &models.User{
					ID:user.ID,
					Username:  user.Username,
					LastName:  user.LastName,
					FirstName: user.FirstName,
					Email: user.Email,
				}, nil
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.User); ok && v.ID == user.ID {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	return authMiddleware, err
}

func InitMiddleware(Conn *gorm.DB) *GoMiddleware  {
	return &GoMiddleware{Conn:Conn}
}