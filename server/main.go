package main

import (
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	id       int
	username string
	password string
}

type Person struct {
	id      int
	name    string
	address string
}

var client *mongo.Client

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	doc := User{id: 1, username: username, password: password}

	if err := InsertOne(doc, client); err != nil {
		fmt.Println(err.Error())
	}

}

func main() {
	// link for the mongoDB
	uri := getUri()

	// connect to the mongoDB
	client = Connect(uri)

	defer Disconnect(client)

	// define static server
	fileServer := http.FileServer(http.Dir("../static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	fmt.Printf("Starting server at port 8080\n")

	// Listen and serve port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
