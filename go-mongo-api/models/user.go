package models

import (
	"context"

	"github.com/Naveenchand06/go-projects/go-mongo-api/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string 		  `json:"name" bson:"name"`
	Email string 		  `json:"email" bson:"email"`
	MobileNumber string   `json:"mobileNumber" bson:"mobileNumber"`
}

func (u *User) CreateUser(db *mongo.Client) error {
	usersCollection := db.Database(constants.DBName).Collection(constants.UserCollection)
	result, err := usersCollection.InsertOne(context.TODO(), u)
	if err != nil {
		return err
	}
	u.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}