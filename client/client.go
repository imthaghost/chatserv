package client

import "github.com/imthaghost/chatserv/protocol"

type messageHandler func(string)

// ChatClient ....
type ChatClient interface {
	Dial(address string) error
	Start()
	Close()
	Send(command interface{}) error
	SetName(name string) error
	SendMessage(message string) error
	Incoming() chan protocol.MessageCommand
}
