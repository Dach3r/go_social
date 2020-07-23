package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*User is the struct table for MongoDB*/
type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FullName    string             `bson:" fullName" json:"fullName,omitempty"`
	DateOfBirth time.Time          `bson:" dataOfBirth" json:"dataOfBirth,omitempty"`
	Email       string             `bson:" email" json:"email"`
	Password    string             `bson:" password" json:"password,omitempty"`
	Avatar      string             `bson:" avatar" json:"avatar"`
}
