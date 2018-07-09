package auth_controller

import (
	"fmt"

	"../models"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	if err := auth_model.CreateUser(c); err != nil {
		fmt.Println(err)
	}
	return
}

func JwtHandler(c *gin.Context) {
	auth_model.AuthJwt(c)
}
