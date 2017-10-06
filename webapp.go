package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mekicha/telebot"
)

func main() {
	//  mux := http.NewServeMux()
	bot, err := telebot.NewBot(os.Getenv("TOKEN"))

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Owner.Username)

	updates := bot.ListenForWebhook("/" + bot.Token)

for update := range updates {
		log.Printf("%+v\n", update)
	}
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
	
}

// func handler(res http.ResponseWriter, req *http.Request) {
	
// }

// func botHandler(res http.ResponseWriter, req *http.Request) {
// 	var update telebot.Update

// 	decoder := json.NewDecoder(req.Body)

// 	err := decoder.Decode(&update)
// 	if err != nil {
// 		log.Println( "Error decoding request body")
// 	}
// 	fmt.Fprintf(res, update)
// }
