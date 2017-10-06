package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mekicha/telebot"
)

func main() {
	//  mux := http.NewServeMux()
	 http.HandleFunc("/", handler)

	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
	
}

func handler(res http.ResponseWriter, req *http.Request) {
	bot, err := telebot.NewBot(os.Getenv("TOKEN"))

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Owner.Username)

	updates, _ := bot.ListenForWebhook("/" + bot.Token)

	for update := range updates {
		log.Println(update)
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
