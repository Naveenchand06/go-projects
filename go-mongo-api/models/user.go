package models

import (
	"context"
	"errors"

	"github.com/Naveenchand06/go-projects/go-mongo-api/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string 		  `json:"name,omitempty" bson:"name,omitempty"`
	Email string 		  `json:"email,omitempty" bson:"email,omitempty"`
	MobileNumber string   `json:"mobileNumber,omitempty" bson:"mobileNumber,omitempty"`
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

func GetUserById(db *mongo.Client, id string) (*User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if(err != nil) {
		return nil, errors.New("not a valid id")
	}
	var user User
	usersCollection := db.Database(constants.DBName).Collection(constants.UserCollection)
	err = usersCollection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) UpdateUserById(db *mongo.Client, id string) (*User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid documnet id")
	}
	u.Id = objectId
	updateDoc := bson.M{}
	if u.Name != "" {
		updateDoc["name"] = u.Name
	}
	if u.Email != "" {
		updateDoc["email"] = u.Email
	}
	if u.MobileNumber != "" {
		updateDoc["mobileNumber"] = u.MobileNumber
	}
	filter := bson.M{"_id": objectId}
	update := bson.M{
		"$set": updateDoc,
	}
	usersCollection := db.Database(constants.DBName).Collection(constants.UserCollection)
	_, err = usersCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	res := usersCollection.FindOne(context.TODO(), filter)
	res.Decode(&u)	
	return u, nil;
}

func DeleteUserByID(db *mongo.Client, id string) (string, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", errors.New("not a valid id")
	}

	usersCollection := db.Database(constants.DBName).Collection(constants.UserCollection)
	result, err := usersCollection.DeleteOne(context.TODO(), bson.M{"_id": objectId})

	if err != nil {
		return "", err
	}
	if result.DeletedCount > 0 {
		return "document deleted successfully", nil
	} else {
		return "document not found", errors.New("documnet not found")
	}
}