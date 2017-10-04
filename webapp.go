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
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, this is our first Go web app on Heroku!")
}
