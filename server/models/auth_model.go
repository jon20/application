package auth_model

import (
	"fmt"
	"log"
	"time"

	"../db"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Serv struct {
	Server   string
	Database string
}

var db *mgo.Database

const COLLECTION = "users"

func Connect() {
	var m = Serv{}

	m.Server = "localhost"
	m.Database = "user"
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(m.Database)

}

func CreateUser(c *gin.Context) error {

	var use = user.User{}
	if err := c.BindJSON(&use); err != nil {
		fmt.Println("error")
	}
	use.Created_at = bson.Now()
	use.Updated_at = bson.Now()

	if err := Insert(use); err != nil {
		return err
	}
	return nil
}

func Insert(user user.User) error {
	if err := db.C(COLLECTION).Insert(&user); err != nil {
		fmt.Println("err")
		return err
	}
	return nil

}
func FindByName(name string) error {
	var use = user.User{}
	fmt.Println(name)
	err := db.C(COLLECTION).Find(bson.M{"name": name}).One(&use)
	return err
}

func AuthJwt(c *gin.Context) {
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: ValidateAuth,
		Unauthorized:  Unauth,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
	authMiddleware.LoginHandler(c)
}

func Unauth(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func ValidateAuth(us string, pass string, c *gin.Context) (interface{}, bool) {
	var use = user.User{}
	//if err := c.BindJSON(&use); err != nil {
	//	fmt.Println("error")
	//}
	fmt.Println(us)
	err := FindByName(us)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	fmt.Println("Find")
	return use, true

}
