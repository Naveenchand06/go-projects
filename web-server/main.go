package main

import (
	"fmt"
	"log"
	"net/http"
)


func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Welcome to hello route", )

}


func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse Form err: %v", err)
		return
	}

	fmt.Fprintln(w, "POST Request successfull")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "The name is: %v \n The address is: %v", name, address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", homeHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("starting server at port 8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}