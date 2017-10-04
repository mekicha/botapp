package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/mekicha/telebot"
)

func main() {

	bot, err := telebot.NewBot(os.Getenv("TOKEN"))

	if err != nil {
		log.Fatal(err)
	}
	updates := bot.ListenForWebhook("/" + bot.Token)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
	http.HandleFunc("/", handler)
	fmt.Println("Listening...")
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, this is our first Go web app on Heroku!")
}
