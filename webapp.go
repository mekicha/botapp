package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/mekicha/telebot"
)

func main() {
	
	http.HandleFunc("/", handler)
	
	// http.HandleFunc("/" + bot.Token, botHandler)
	fmt.Println("Listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("Listen and Serve Error:")
	}

}

func handler(res http.ResponseWriter, req *http.Request) {
	bot, err := telebot.NewBot(os.Getenv("TOKEN"))

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Owner.Username)
	updates := bot.ListenForWebhook("/" + bot.Token)
	
	for update := range updates {
		fmt.Fprintf(res, "%+v\n", update)
	}
}

// func botHandler(res http.ResponseWriter, req *http.Request) {
// 	var update telebot.Update

// 	decoder := json.NewDecoder(req.Body)

// 	err := decoder.Decode(&update)
// 	if err != nil {
// 		log.Println( "Error decoding request body")
// 	}
// 	fmt.Fprintf(res, update)
// }
