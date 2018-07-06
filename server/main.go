package main

import (

	//jwt "github.com/dgrijalva/jwt-go"
	//"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-contrib/cors"

	"./controllers"

	"github.com/gin-gonic/gin"
)

var secretKey = "75c92a074c341e9964329c0550c2673730ed8479c885c43122c90a2843177d5ef21cb50cfadcccb20aeb730487c11e09ee4dbbb02387242ef264e74cbee97213"

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	router.Use(cors.New(config))
	router.POST("/auth", auth_controller.SignIn)

	router.Run(":8000")
}
