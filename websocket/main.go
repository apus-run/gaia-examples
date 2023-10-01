package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/apus-run/gaia"
	"github.com/apus-run/gaia/transport/websocket"
	"github.com/apus-run/sea-kit/encoding"
)

var testServer *websocket.Server

const (
	MessageTypeChat = iota + 1
)

type ChatMessage struct {
	Type    int    `json:"type"`
	Message string `json:"message"`
}

func main() {
	wsSrv := websocket.NewServer(
		websocket.Address(":8800"),
		websocket.Path("/ws"),
		websocket.ConnectHandle(handleConnect),
		websocket.Codec(encoding.GetCodec("json")),
	)

	testServer = wsSrv

	wsSrv.RegisterMessageHandler(MessageTypeChat,
		func(sessionId websocket.SessionID, payload websocket.MessagePayload) error {
			switch t := payload.(type) {
			case *ChatMessage:
				return handleChatMessage(sessionId, t)
			default:
				return errors.New("invalid payload type")
			}
		},
		func() websocket.Any { return &ChatMessage{} },
	)

	app := gaia.New(
		gaia.WithName("websocket"),
		gaia.WithServer(
			wsSrv,
		),
	)
	if err := app.Run(); err != nil {
		log.Println(err)
	}
}

func handleConnect(sessionId websocket.SessionID, register bool) {
	if register {
		fmt.Printf("%s connected\n", sessionId)
	} else {
		fmt.Printf("%s disconnect\n", sessionId)
	}
}

func handleChatMessage(sessionId websocket.SessionID, message *ChatMessage) error {
	fmt.Printf("[%s] Payload: %v\n", sessionId, message)

	testServer.Broadcast(MessageTypeChat, *message)

	return nil
}
