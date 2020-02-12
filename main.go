package main

import "fmt"

func main() {
	// var (
	// 	port      = os.Getenv("PORT")       // set automatically
	// 	publicURL = os.Getenv("PUBLIC_URL") // from config vars
	// 	token     = os.Getenv("TOKEN")      // from config vars
	// )

	// webhook := &tb.Webhook{
	// 	Listen:   ":" + port,
	// 	Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	// }

	// pref := tb.Settings{
	// 	Token:  token,
	// 	Poller: webhook,
	// }

	// b, err := tb.NewBot(pref)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// b.Handle("/hello", func(m *tb.Message) {
	// 	b.Send(m.Sender, "Hi!")
	// })
	fmt.Println("Hello !")
}
