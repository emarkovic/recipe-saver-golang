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

/*
recipe randomizer
- save recipe url
- delete recipe url
- favorite recipe
- assign recipe for a day of the week? -> how?
*/

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
	// http.HandleFunc("/v1/recipes", recipesHander)

	addr := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("server is listening at %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
