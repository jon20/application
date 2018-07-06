package auth_controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	fmt.Println(c.PostForm("text"))
	c.JSON(200, gin.H{"message": c.PostForm("text")})
}
