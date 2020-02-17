package main

import (
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		// port      = os.Getenv("PORT")       // set automatically
		// publicURL = os.Getenv("PUBLIC_URL") // from config vars
		token = os.Getenv("TOKEN") // from config vars
	)

	log.Println(token)

	poller := &tb.LongPoller{
		Timeout: 10 * time.Second,
	}

	pref := tb.Settings{
		// URL:    publicURL,
		Token:  "891489265:AAHt-P6qMWV5n8gPH88_xdeFsahm-hEU7k0",
		Poller: poller,
	}

	b, err := tb.NewBot(pref)

	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hi!")
	})

	b.Start()
}
