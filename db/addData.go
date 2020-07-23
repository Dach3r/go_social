package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/dach3r/vecker/models"
	"context"
	"time"
)

/*InsertRegister this function insert a new register user in db*/
func InsertRegister(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("vecker")
	column := db.Collection("user")
	
	user.Password, _ = EncryptPassword(user.Password)

	result, err := column.InsertOne
}

/*VerifyUserExist check if user exis on db*/
func VerifyUserExist() {

}
