package main

import (
	"net/http"

	"github.com/Naveenchand06/go-projects/go-mongo/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)


func main() {
	router := httprouter.New()
	controller := controllers.NewUserController(getSession())
	router.GET("/user/:id", controller.GetUser)
	router.POST("/user", controller.CreateUser)
	router.DELETE("/user/:id", controller.DeleteUser)

	// * Listen for incoming requests
	http.ListenAndServe(":5010", router)

}


func getSession() *mgo.Session {
	session, err :=  mgo.Dial("mongodb://myuser:mypass@localhost:27017/mongo-go")
	if err != nil {
		panic(err)
	}
	return session
}