package commands

import (
	"bot/discord"
	"bot/discord/messages"
	"bot/utils/builders"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const wolframApiUrlTemplate = "http://api.wolframalpha.com/v1/conversation.jsp?i=%s&appid=%s"

type WolframResponse struct {
	Result         string `json:"result"`
	ConversationID string `json:"conversationID"`
	Host           string `json:"host"`
	S              string `json:"s"`
}

type WolframCommand struct {
	interactionCreate messages.InteractionCreate
}

func (d WolframCommand) Execute() error {

	response, err := http.Get(fmt.Sprintf(wolframApiUrlTemplate, url.QueryEscape(d.interactionCreate.D.Data.Options[0].Value.(string)), url.QueryEscape(os.Getenv("WOLFRAM_APPLICATION_ID"))))
	if err != nil {
		return fmt.Errorf("wolfram api get: %w", err)
	}

	bytesResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("wolfram api response read: %w", err)
	}
	var wolframResponse WolframResponse
	json.Unmarshal(bytesResponse, &wolframResponse)

	reply := "I don't know bro."
	if wolframResponse.Result != "" {
		reply = wolframResponse.Result
	}

	interactionCallback := builders.CreateInteractionCallback().
		AddContent(
			fmt.Sprintf("%s: %s", d.interactionCreate.D.Data.Options[0].Value, reply)).
		Get()

	discord.PostInteractionResponse(d.interactionCreate.D.ID, d.interactionCreate.D.Token, interactionCallback)

	return nil
}

func (d WolframCommand) Respond() error {
	return nil
}
