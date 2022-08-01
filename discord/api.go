package discord

import (
	"bot/discord/messages"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/net/websocket"
)

const discordSocketUrl = "wss://gateway.discord.gg/?v=10&encoding=json"
const discordCallbackTemplateUrl = "https://discord.com/api/v10/interactions/%s/%s/callback"
const discordGetCallbackTemplateUrl = "https://discord.com/api/v10/webhooks/%s/%s/messages/@original"
const discordEditCallbackTemplateUrl = "https://discord.com/api/v10/webhooks/%s/%s/messages/@original"
const discordFollowUpTemplateUrl = "https://discord.com/api/v10/webhooks/%s/%s/"
const discordPostChannelBotInfoUrl = "https://discord.com/api/v10/channels/1003011209848701068/messages"

func Identify(conn *websocket.Conn, appId string) {

	var identifyPayload messages.Identify
	identifyPayload.D.Token = appId
	identifyPayload.Op = 2
	identifyPayload.D.Intents = int(int8(1) << int8(4))
	identifyPayload.D.Properties.Browser = "placeholder"
	identifyPayload.D.Properties.Device = "placeholder"
	identifyPayload.D.Properties.Os = "linux"

	if err := websocket.JSON.Send(conn, identifyPayload); err != nil {
		log.Print(err)
	}
}

func Heartbeat(heartbeat int, conn *websocket.Conn, errorChan chan error) error {

	ticker := time.NewTicker(time.Duration(float32(heartbeat)*rand.Float32()) * time.Millisecond)
	done := make(chan bool)

	heartBeatPayload := messages.HeartBeat{Op: 1, D: 0}
	if err := websocket.JSON.Send(conn, heartBeatPayload); err != nil {
		return fmt.Errorf("sending heartbeat: %w", err)
	}

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				log.Printf("debug sending heartbeat\n\n")
				if err := websocket.JSON.Send(conn, heartBeatPayload); err != nil {
					errorChan <- fmt.Errorf("sending heartbeat: %w", err)
					return
				}
			}
		}
	}()

	return nil
}

func Connect() (*websocket.Conn, int, error) {

	conn, err := websocket.Dial(discordSocketUrl, "", "http://localhost")
	if err != nil {
		return nil, 0, fmt.Errorf("connect websocket: %w", err)
	}

	var helloPayload messages.Hello
	if err := websocket.JSON.Receive(conn, &helloPayload); err != nil {
		return nil, 0, fmt.Errorf("recieve hello: %w", err)
	}

	return conn, helloPayload.D.HeartbeatInterval, nil
}

func Listen(conn *websocket.Conn, callback func([]byte) error, errorChan chan error) {

	for {
		var message []byte
		if err := websocket.Message.Receive(conn, &message); err != nil {
			errorChan <- fmt.Errorf("websocket recieve: %w", err)
			return
		}
		log.Printf("debug raw_message: %s\n\n", message)

		var payload messages.Base
		json.Unmarshal(message, &payload)

		switch payload.T {
		case "INTERACTION_CREATE":
			go interactionCreateRoutine(message, callback)
		default:
		}
	}
}

func interactionCreateRoutine(message []byte, callback func([]byte) error) {

	fmt.Printf("INTERACTION_CREATE\n\n")
	if err := callback(message); err != nil {
		log.Print(err)
	}
}
