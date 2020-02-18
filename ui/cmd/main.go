package main

import (
	"flag"
	"log"

	"github.com/imthaghost/gochat/client"
	"github.com/imthaghost/gochat/ui"
)

func main() {
	address := flag.String("server", "localhost:3333", "Which server to connect to")

	flag.Parse()

	client := client.NewClient()
	err := client.Dial(*address)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// start the client to listen for incoming message
	go client.Start()

	ui.StartUi(client)
}
