package discord

import (
	"bot/discord/messages"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func PostChannelMessage(channelMessage messages.ChannelMessage) error {

	channelMessagePayload, err := json.Marshal(channelMessage)
	if err != nil {
		return fmt.Errorf("marshaling channel message: %s", err)
	}

	request, err := http.NewRequest("POST", discordPostChannelBotInfoUrl, bytes.NewBuffer(channelMessagePayload))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("DISCORD_APPLICATION_ID")))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("post channel message: %w", err)
	}

	body, _ := io.ReadAll(response.Body)
	fmt.Printf("%s\n\n", string(body))

	return nil
}
