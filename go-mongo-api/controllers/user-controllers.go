package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Naveenchand06/go-projects/go-mongo-api/config"
	"github.com/Naveenchand06/go-projects/go-mongo-api/constants"
	"github.com/Naveenchand06/go-projects/go-mongo-api/models"
)

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func sendError(w http.ResponseWriter, statusCode int, err string) {
	error_message := map[string]interface{}{"error": err}
	sendResponse(w, statusCode, error_message)
}

// *************** Controller Functions *********************

func CreateUserController(w http.ResponseWriter, req *http.Request) {
	var reqUser models.User
	err := json.NewDecoder(req.Body).Decode(&reqUser)
	if err != nil {
		fmt.Println("The error is ->", err)
		sendError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	err = reqUser.CreateUser(config.GetDB())

	if err != nil {
		sendError(w, http.StatusInternalServerError, "something went wrong")
		return
	}
	sendResponse(w, http.StatusOK, reqUser)
}

func GetAllUserController(w http.ResponseWriter, request *http.Request) {
	usersCollection := config.GetDB().Database(constants.DBName).Collection(constants.UserCollection)
	cursor, err := usersCollection.Find(context.TODO(), map[string]interface{}{})
	if err != nil {
		fmt.Println("The error is ->", err)
		sendError(w, http.StatusInternalServerError, "something went wrong 1")
		return
	}
	users := make([]models.User, 0)
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			sendError(w, http.StatusInternalServerError, "something went wrong 2")
			return
		}
		users = append(users, user)
	}
	sendResponse(w, http.StatusOK, users)
}

func GetUserByIDController(writer http.ResponseWriter, request *http.Request) {

}

func UpdateUserController(writer http.ResponseWriter, request *http.Request) {

}

func DeleteUserController(writer http.ResponseWriter, request *http.Request) {

}