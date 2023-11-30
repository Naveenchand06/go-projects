package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)


type App struct {
	Router *mux.Router
	DB *sql.DB
}


func (app *App) Initialise(dbUser string, dbPass string, dbName string) error {
	fmt.Println("Intialise ----> 1 : ", DbName, DbPassword, DbUser)

	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", dbUser, dbPass, dbName)
	fmt.Println("Intialise ----> 1")

	var err error
	app.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("Intialise ----> 3", err)

		return err
	}
	fmt.Println("Intialise ----> 4")

	app.Router = mux.NewRouter().StrictSlash(true)
	fmt.Println("Intialise ----> 5")

	app.handleRoutes()
	return nil
}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}

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

func (app *App) getProducts(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Executed ----> 1")
	products, err := getProducts(app.DB)
	fmt.Println("Executed ----> 2")

	if err != nil {
	fmt.Println("Executed ----> 3", err)

		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("Executed ----> 4")
	sendResponse(w, http.StatusOK, products)
}

func (app * App) getProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get P ---> 1")
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	fmt.Println("get P ---> 2")

	if err != nil {
	fmt.Println("get P ---> 3")

		sendError(w, http.StatusBadRequest, "invalid product id")
		return
	}
	fmt.Println("get P ---> 4")

	p :=  product{ID: key}
	fmt.Println("get P ---> 5")

	err = p.getProduct(app.DB)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			sendError(w, http.StatusNotFound, err.Error())
		default:
			sendError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	sendResponse(w, http.StatusOK, p)
}

func (app * App) createProduct(w http.ResponseWriter, r *http.Request) {
	var p product
	err := json.NewDecoder(r.Body).Decode(&p);
	if(err != nil) {
		sendError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	err = p.createProduct(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusCreated, p)
}

func (app *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid path variable")
		return
	}
	var p product
	err = json.NewDecoder(r.Body).Decode(&p)
	p.ID = key
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err = p.updateProduct(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(w, http.StatusOK, p)
}


func (app *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid path variable")
		return
	}
	p := product{ID: key}
	err = p.deleteProduct(app.DB) 
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, map[string]string{"message": "Deleted successfully"})
}

func (app *App) handleRoutes() {
	fmt.Println("Handle ----> 1")

	app.Router.HandleFunc("/products", app.getProducts).Methods("GET")
	app.Router.HandleFunc("/product/{id}", app.getProduct).Methods("GET")
	app.Router.HandleFunc("/product", app.createProduct).Methods("POST")
	app.Router.HandleFunc("/product/{id}", app.updateProduct).Methods("PUT")
	app.Router.HandleFunc("/product/{id}", app.deleteProduct).Methods("DELETE")
}