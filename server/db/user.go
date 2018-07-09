package user

import "time"

type User struct {
	ID         int       `bson:"id" json:"id"`
	Name       string    `bson:"name" json:"name"`
	Email      string    `bson:"email" json:"email"`
	Password   string    `bson:"password" json:"password"`
	Created_at time.Time `bson:"created_at" json:"created_at"`
	Updated_at time.Time `bson:"updated_at" json:"updated_at"`
}
