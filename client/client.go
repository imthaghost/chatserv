package client

import "github.com/imthaghost/gochat/protocol"

type ChatClient interface {
	Dial(address string) error
	Start()
	Close()
	Send(command interface{}) error
	SetName(name string) error
	SendMessage(message string) error
	Error() chan error
	Incoming() chan protocol.MessageCommand
}
