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
		Token:  token,
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
