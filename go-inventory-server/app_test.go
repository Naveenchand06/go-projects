package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App
func TestMain(m *testing.M) {
	err := a.Initialise(DbUser, DbPassword, "test_inventory")
	if err != nil {
		log.Fatal("Error occured while initializing database")
	}
	createTable()
	m.Run()
}

func createTable() {
	createTableQuery := `CREATE TABLE IF NOT EXISTS products (
		id int NOT NULL AUTO_INCREMENT,
		name varchar(255) NOT NULL,
		quantity int,
		price float(10,7),
		PRIMARY KEY (id)
		);`
	_, err := a.DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("Delete from products")
	a.DB.Exec("ALTER TABLE products AUTO_INCREMENT=1")
}

func addProduct(name string, qunatity int, price float64) {
	query := fmt.Sprintf("insert into products (name, quantity, price) values('%v', %v, %v)", name, qunatity, price)
	_, err := a.DB.Exec(query)
	if err != nil {
		log.Println(err)
	}
}

func checkStatusCode(t *testing.T, expectedStatusCode int, actualStatusCode int) {
	if expectedStatusCode != actualStatusCode {
		t.Errorf("Expected Status: %v, Received Status: %v", expectedStatusCode, actualStatusCode)
	}
}

func sendRequest(request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, request)
	return recorder
}

func TestGetProduct(t *testing.T) {
	clearTable()
	addProduct("iMAC", 3, 2)
	request, _ := http.NewRequest("GET", "/product/1", nil)
	response := sendRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)
}


func TestCreateProduct(t *testing.T) {
	var product = []byte(`{"name": "Chair", "quantity": 3, "price": 100}`)
	request, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(product))
	request.Header.Set("Content-type", "application/json")
	response := sendRequest(request)
	checkStatusCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "Chair" {
		t.Errorf("Expected is: %v, but got %v", "Chair", m["name"])
	}

	if m["quantity"] != 3.0 {
		t.Errorf("Expected is: %v, but got %v", 3.0, m["quantity"])
	}
}
func TestDeleteProduct(t *testing.T) {
	clearTable()
	addProduct("Bed", 10, 10)
	request, _ := http.NewRequest("DELETE", "/product/1", nil)
	response := sendRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)

	request, _ = http.NewRequest("GET", "/product/1", nil)
	response = sendRequest(request)
	checkStatusCode(t, http.StatusNotFound, response.Code)
}

func TestUpdateProduct(t *testing.T) {
	clearTable()
	addProduct("sofa", 10, 10)
	request, _ := http.NewRequest("GET", "/product/1", nil)
	response := sendRequest(request)

	var oldValue map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &oldValue)
	log.Println("New: -> ", oldValue)

 
	var product = []byte(`{"name": "sofa", "quantity": 1, "price": 1}`)
	request, _ = http.NewRequest("PUT", "/product/1", bytes.NewBuffer(product))
	request.Header.Set("Content-type", "application/json")
	response = sendRequest(request)

	var newValue map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &newValue)
	log.Println(newValue)

	if oldValue["id"] != newValue["id"] {
		t.Errorf("Expected id: %v, Got %v", oldValue["id"], newValue["id"])
	}

	if oldValue["price"] == newValue["price"] {
		t.Errorf("Expected is: %v, Got %v", oldValue["price"], newValue["price"])
	}

}