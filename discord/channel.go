package discord

import (
	"bot/discord/messages"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func PostChannelMessage(channelMessage messages.ChannelMessage) error {

	channelMessagePayload, err := json.Marshal(channelMessage)
	if err != nil {
		return fmt.Errorf("error marshaling channel message: %s", err)
	}

	request, err := http.NewRequest("POST", discordPostChannelBotInfoUrl, bytes.NewBuffer(channelMessagePayload))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bot OTkyNTA4MDg5NDYwODYzMDM3.G5BwZ6.lJHFJWmzTQPGYE3bjQZoE_mW9zXoOFUuUeQhRk")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("post channel message: %w", err)
	}

	body, _ := io.ReadAll(response.Body)
	fmt.Printf("%s\n\n", string(body))

	return nil
}
