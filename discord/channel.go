package discord

import (
	"bot/discord/messages"
	"fmt"
)

func PostChannelMessage(channelID string, channelMessage messages.ChannelMessage) error {

	_, _, err := client.DoPostObject(getMessageChannelEndpointById(channelID), channelMessage)
	if err != nil {
		return fmt.Errorf("post object: %w", err)
	}

	return nil
}
