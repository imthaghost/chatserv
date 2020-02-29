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
<!-- ![landing](docs/media/chat.png) -->

### 📚 Table of Contents

1. [Deliverables](#deliverables)
2. [Getting Started](#getting-started)
3. [ProjectStructure](#project-structure)

## Deliverables

-   [x] Single chat room
-   [x] Users can connect to the server
-   [x] Users can set their name
-   [x] Users can send the message to the room
-   [x] Users can see all other user's messages

## Getting started

```bash
# in the server directory start the tcp-server
go run main.go
# change to user-interface directory
cd tui
# dial to the tcp-sever
go run main.go --server localhost:3333

```

## Project Structure

```bash
📂 chatserv
├── README.md
├── client
│   ├── client.go
│   └── tcp_client.go
├── docs
│   └── media
│       └── chat.png
├── go.mod
├── go.sum
├── protocol
│   ├── command.go
│   ├── reader.go
│   ├── reader_test.go
│   ├── writer.go
│   └── writer_test.go
├── server
│   ├── cmd
│   │   └── main.go
│   ├── server.go
│   └── tcp_server.go
└── tui
    ├── chatview.go
    ├── cmd
    │   └── main.go
    ├── loginview.go
    └── tui.go
```
