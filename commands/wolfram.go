package commands

import (
	"bot/discord"
	"bot/discord/messages"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type WolframResponse struct {
	Result         string `json:"result"`
	ConversationID string `json:"conversationID"`
	Host           string `json:"host"`
	S              string `json:"s"`
}

func wolframCommand(interactionCreate messages.InteractionCreate) error {

	response, err := http.Get(fmt.Sprintf("http://api.wolframalpha.com/v1/conversation.jsp?i=%s&appid=8PHTWK-KL2R5P6WEU", url.QueryEscape(interactionCreate.D.Data.Options[0].Value.(string))))
	if err != nil {
		return fmt.Errorf("wolfram command: %s", err)
	}

	bytesResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("wolfram command: %s", err)
	}
	var wolframResponse WolframResponse
	json.Unmarshal(bytesResponse, &wolframResponse)

	var reply string
	if wolframResponse.Result == "" {
		reply = "I don't know bro."
	} else {
		reply = wolframResponse.Result
	}

	interactionCallback := messages.InteractionCallback{
		Type: 4,
		Data: messages.Data{
			Content: fmt.Sprintf("%s: %s", interactionCreate.D.Data.Options[0].Value, reply),
		},
	}

	discord.PostInteractionCallback(interactionCreate.D.ID, interactionCreate.D.Token, &interactionCallback)
	return nil
}
