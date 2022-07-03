package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

const discordSocketUrl = "wss://gateway.discord.gg/?v=10&encoding=json"
const discordCallbackTemplateUrl = "https://discord.com/api/v10/interactions/%s/%s/callback"

func postInteractionCallback(id string, token string, interactionCallbackPayload *InteractionCallbackPayload) error {

	callbackPayload, err := json.Marshal(interactionCallbackPayload)
	if err != nil {
		return fmt.Errorf("error marshaling callback: %s", err)
	}

	callback := fmt.Sprintf(discordCallbackTemplateUrl, id, token)
	if _, err := http.Post(callback, "application/json", bytes.NewBuffer(callbackPayload)); err != nil {
		return fmt.Errorf("error during post to callback: %s", err)
	}

	return nil
}

func Identify(conn *websocket.Conn) {

	var identifyPayload IdentifyPayload
	identifyPayload.D.Token = "OTkyNTA4MDg5NDYwODYzMDM3.G5BwZ6.lJHFJWmzTQPGYE3bjQZoE_mW9zXoOFUuUeQhRk"
	identifyPayload.Op = 2
	identifyPayload.D.Intents = int(int8(1) << int8(4))
	identifyPayload.D.Properties.Browser = "placeholder"
	identifyPayload.D.Properties.Device = "placeholder"
	identifyPayload.D.Properties.Os = "linux"

	if err := websocket.JSON.Send(conn, identifyPayload); err != nil {
		fmt.Println(err)
	}
}

func Heartbeat(heartbeat int, conn *websocket.Conn) {

	ticker := time.NewTicker(time.Duration(float32(heartbeat)*rand.Float32()) * time.Millisecond)
	done := make(chan bool)

	heartBeatPayload := HeartBeatPayload{Op: 1, D: 0}
	if err := websocket.JSON.Send(conn, heartBeatPayload); err != nil {
		fmt.Println(err)
	}

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				fmt.Printf("debug sending heartbeat\n")
				if err := websocket.JSON.Send(conn, heartBeatPayload); err != nil {
					fmt.Println(err)
				}
			}
		}
	}()
}

func Connect() (*websocket.Conn, int) {

	conn, err := websocket.Dial(discordSocketUrl, "", "http://localhost")
	if err != nil {
		fmt.Println(err)
	}

	var helloPayload HelloPayload
	if err := websocket.JSON.Receive(conn, &helloPayload); err != nil {
		fmt.Println(err)
	}

	return conn, helloPayload.D.HeartbeatInterval
}

func Listen(conn *websocket.Conn) {

	for {
		var message []byte
		if err := websocket.Message.Receive(conn, &message); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("debug raw_message: %s\n\n", message)

		var payload Payload
		json.Unmarshal(message, &payload)

		switch payload.T {

		case "INTERACTION_CREATE":
			fmt.Printf("INTERACTION_CREATE\n")
			if err := handleInteraction(message); err != nil {
				fmt.Println(err)
			}

		case "GUILD_CREATE":
			var guildCreatePayload GuildCreatePayload
			json.Unmarshal(message, &guildCreatePayload)
			fmt.Printf("GUILD_CREATE, %s, %s\n", guildCreatePayload.D.JoinedAt, guildCreatePayload.D.Name)

		case "MESSAGE_CREATE":
			var messageCreatePayload MessageCreatePayload
			json.Unmarshal(message, &messageCreatePayload)
			fmt.Printf("MESSAGE_CREATE, %s, %s, %s\n", messageCreatePayload.D.Timestamp, messageCreatePayload.D.Author.Username, messageCreatePayload.D.Content)

		default:
		}
	}
}

func handleInteraction(message []byte) error {

	command, interactionCreatePayload := readInteractionCreatePayload(message)

	switch command {

	case "hello":
		helloCommand(interactionCreatePayload)

	case "pirate":
		pirateCommand(interactionCreatePayload)

	case "dice":
		diceCommand(interactionCreatePayload)

	}

	return nil
}
