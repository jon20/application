package main

import (

	//jwt "github.com/dgrijalva/jwt-go"
	//"github.com/dgrijalva/jwt-go/request"

	"github.com/gin-contrib/cors"

	"./controllers"
	"./models"
	"github.com/gin-gonic/gin"
)

func init() {
	auth_model.Connect()
}

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func main() {

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	router.Use(cors.New(config))
	router.POST("/auth", auth_controller.SignIn)
	router.POST("/login", auth_controller.JwtHandler)
	router.Run(":8000")
}
