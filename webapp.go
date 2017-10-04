package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening...")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, this is our first Go web app on Heroku!")
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT env var detected, defaulting to " + port)
	}
	return ":" + port 
}