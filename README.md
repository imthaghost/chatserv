<p align="center">
  <!-- <a href="#">
    <img alt="jive-search logo" src="https://github.com/imthaghost/makescraper/blob/master/docs/media/logo.jpg"> 
  </a> -->
</p>

<br>

<p align="center">
A chatserver written in GO 
</p>

<br>
<p align="center">
   <a href="#">
    <img src="https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg"alt="gitmoji-changelog">
  </a>
  <a href="https://goreportcard.com/badge/github.com/imthaghost/gochat"><img src="https://goreportcard.com/badge/github.com/imthaghost/gochat"></a>
</p>
<br>

<br>
<p align="center">
   <a href="#"><img src="https://github.com/imthaghost/gochat/blob/master/docs/media/chat.png"></a>  
</p>
<br>

### 📚 Table of Contents

1. [Deliverables](#deliverables)
2. [Getting Started](#getting-started)
3. [Local Development](#local-development)

## Deliverables

-   [x] Single chat room
-   [x] Users can connect to the server
-   [x] Users can set their name
-   [x] Users can send the message to the room
-   [x] Users can see all other user's messages

## Installation

```sh
$ go get github.com/imthaghost/chatserv
```

## Getting Started

### server.go

```go
package main

import (
	"github.com/imthaghost/chatserv/server"
)

func main() {
	var s server.ChatServer
	s = server.NewServer()
	s.Listen(":3333")

	// start the server
	s.Start()
}
```

### client.go

```go
package main

import (
	"flag"
	"log"

	"github.com/imthaghost/chatserv/client"
	"github.com/imthaghost/chatserv/tui"
)

func main() {
	address := flag.String("server", "", "Which server to connect to")

	flag.Parse()

	client := client.NewClient()
	err := client.Dial(*address)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// start the client to listen for incoming message
	go client.Start()

	tui.StartUi(client)
}
```

### Run

```sh
# we start the irc server
$ go run server.go
# in a new termnal window we start the client
$ go run client.go --server localhost:3333
```

![Example](/docs/media/client.gif)

## Local Development

```sh
# inside github.com/imthaghost/server/ircserver
$ go run main.go
# inside github.com/imthaghost/tui/irc
$ go run main.go --server localhost:3333
```
