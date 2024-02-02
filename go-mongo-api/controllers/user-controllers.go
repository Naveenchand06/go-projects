package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Naveenchand06/go-projects/go-mongo-api/config"
	"github.com/Naveenchand06/go-projects/go-mongo-api/constants"
	"github.com/Naveenchand06/go-projects/go-mongo-api/models"
	"github.com/Naveenchand06/go-projects/go-mongo-api/utils"
	"github.com/gorilla/mux"
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

// ***** Create User *****
func CreateUserController(w http.ResponseWriter, req *http.Request) {
	var reqUser models.User
	// err := json.NewDecoder(req.Body).Decode(&reqUser)
	err := utils.DecodeRequestBody(req, &reqUser)
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

// ***** Get All User *****
func GetAllUserController(w http.ResponseWriter, req *http.Request) {
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

func GetUserByIDController(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	user, err := models.GetUserById(config.GetDB() ,id)
	if err != nil {
		sendError(w, http.StatusNotFound, err.Error())
		return
	}
	sendResponse(w, http.StatusFound, user)
}

func UpdateUserController(w http.ResponseWriter, req *http.Request) {

}

func DeleteUserController(w http.ResponseWriter, request *http.Request) {

}