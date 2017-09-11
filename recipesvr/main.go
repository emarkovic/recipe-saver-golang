package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	apiRoot = "/v1"
)

func defaultMsg(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte("Hi, you have made it to the app. I love you, good job."))
}

func recipesHander(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// "saving" -> creating a new recipe
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("You are saving a recipe"))
	case "GET":
		// get all recipe
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("you are getting all recipes"))
	case "PATCH":
		// assign to day of the week
		// favorite
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("you are either assinging a recipe or favoriting it"))
	case "DELETE":
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("you are deleting a recipe"))
	}
}

// create a user
func usersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { // create a new user
		http.Error(w, "invalid request method", http.StatusBadRequest)
		return
	}
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}

	host := os.Getenv("HOST")
	if len(host) == 0 {
		host = "localhost"
	}

	http.HandleFunc("/", defaultMsg)
	http.HandleFunc("/v1/recipes", recipesHander)

	addr := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("server is listening at %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
