package discord

import (
	"bot/discord/messages"
	"bot/utils/http"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var client *http.Client

func Init() {

	client = getClient()
}

func getClient() *http.Client {

	headers := make(map[string]string)

	headers["Authorization"] = fmt.Sprintf("Bot %s", os.Getenv("DISCORD_BEARER_TOKEN"))

	return &http.Client{
		Headers: headers,
	}
}

func Identify(conn *websocket.Conn, appId string) {

	var identifyPayload messages.Identify
	identifyPayload.D.Token = appId
	identifyPayload.Op = 2
	identifyPayload.D.Intents = int(int8(1) << int8(4))
	identifyPayload.D.Properties.Browser = "placeholder"
	identifyPayload.D.Properties.Device = "placeholder"
	identifyPayload.D.Properties.Os = "linux"

	if err := conn.WriteJSON(identifyPayload); err != nil {
		log.Panic(err)
	}

	log.Print("debug identifyPayload: ")
	log.Print(identifyPayload)
}

func Heartbeat(heartbeat int, conn *websocket.Conn, errorChan chan error) error {

	ticker := time.NewTicker(time.Duration(float32(heartbeat)*rand.Float32()) * time.Millisecond)
	done := make(chan bool)

	heartBeatPayload := messages.HeartBeat{Op: 1, D: 0}
	if err := conn.WriteJSON(heartBeatPayload); err != nil {
		log.Print(fmt.Errorf("sending heartbeat: %w", err))
	}

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				log.Printf("debug sending heartbeat\n\n")
				if err := conn.WriteJSON(heartBeatPayload); err != nil {
					errorChan <- fmt.Errorf("sending heartbeat: %w", err)
				}
			}
		}
	}()

	return nil
}

func Connect() (*websocket.Conn, int, error) {

	conn, _, err := websocket.DefaultDialer.Dial(getGatwayEndpoint(), nil)
	if err != nil {
		return nil, 0, fmt.Errorf("connect websocket: %w", err)
	}

	var helloPayload messages.Hello
	if err := conn.ReadJSON(&helloPayload); err != nil {
		return nil, 0, fmt.Errorf("receive hello: %w", err)
	}

	log.Print("debug hello payload: ")
	log.Print(helloPayload)

	return conn, helloPayload.D.HeartbeatInterval, nil
}

func Listen(conn *websocket.Conn, callback func([]byte) error, errorChan chan error) {

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			errorChan <- fmt.Errorf("websocket receive: %w", err)
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
