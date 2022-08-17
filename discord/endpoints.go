package discord

import (
	"fmt"
	"os"
)

func getOriginalMessageEndpointForToken(token string) string {

	return getWebHookEndpointForToken(token) + "/messages/@original"
}

func getOriginalMessageEndpointForTokenById(token string, messageID string) string {

	return getWebHookEndpointForToken(token) + fmt.Sprintf("/messages/%s", messageID)
}

func getMessageChannelEndpointById(channelID string) string {

	return getChannelsEndpoint() + fmt.Sprintf("/%s/messages", channelID)
}

func getInteractionsCallbackEndpoint(token string) string {

	return getInteractionsEndpoint(token) + "/callback"
}

func getInteractionsEndpoint(token string) string {

	return fmt.Sprintf("https://discord.com/api/v10/interactions/%s/%s", os.Getenv("DISCORD_APPLICATION_ID"), token)
}

func getWebHookEndpointForToken(token string) string {

	return fmt.Sprintf("https://discord.com/api/v10/webhooks/%s/%s", os.Getenv("DISCORD_APPLICATION_ID"), token)
}

func getChannelsEndpoint() string {
	return "https://discord.com/api/v10/channels/"
}

func getGatwayEndpoint() string {
	return "wss://gateway.discord.gg/?v=10&encoding=json"
}
