package main

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"learngo/middleware"
	"learngo/models"
	"log"
	"net/http"
	"os"
	_userRepo "learngo/user/repository"
	_userUsecase "learngo/user/usecase"
	_UserHttpDeliver "learngo/user/delivery/http"
	_courseRepo "learngo/course/repository"
	_courseUsecase "learngo/course/usecase"
	_courseHttpDelver "learngo/course/delivery/http"
	"strconv"
	"time"
)
var db *gorm.DB
var err error

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env")
	}

	usernameDb := os.Getenv("DB_USERNAME")
	passwordDb := os.Getenv("DB_USERNAME")
	hostDb := os.Getenv("DB_HOST")
	portDb := os.Getenv("DB_PORT")
	databaseDb := os.Getenv("DB_DATABASE")

	db, err = gorm.Open("mysql", usernameDb+":"+passwordDb+"@tcp("+hostDb+":"+portDb+")/"+databaseDb+"?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		fmt.Println(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	db.AutoMigrate(&models.User{},&models.Course{})

	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	if port == "" {
		port = "8000"
	}

	middleJwt := middleware.InitMiddleware(db)
	authMiddleware, err := middleJwt.JWT()

	if err !=nil{
		fmt.Println(err)
	}

	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.GET("/refresh_token", authMiddleware.RefreshHandler)

	r.Use(authMiddleware.MiddlewareFunc())

	userRepo:= _userRepo.NewMysqlUserRepository(db)

	var timeOut, _ = strconv.Atoi(os.Getenv("TIMEOUT"))

	timeoutContext := time.Duration(timeOut) * time.Second

	userUsecase:=_userUsecase.NewUserUsecase(userRepo,timeoutContext)

	//usersRoute := r.Group("users")

	_UserHttpDeliver.NewUserHandler(r,userUsecase)

	courseRepo:= _courseRepo.NewMysqlCourseRepository(db)
	courseUsecase:= _courseUsecase.NewCourseUsecase(courseRepo)

	_courseHttpDelver.NewCourseHandler(r,courseUsecase)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
