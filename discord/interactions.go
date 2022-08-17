package discord

import (
	"bot/discord/messages"
	"bot/utils/builders"
	"bot/utils/http"
	"encoding/json"
	"fmt"
)

func PostInteractionResponse(id string, token string, interactionCallbackPayload *messages.InteractionCallback) (string, error) {

	_, bodyResponse, err := client.DoPostObject(getInteractionsCallbackEndpoint(token), interactionCallbackPayload)
	if err != nil {
		return "", fmt.Errorf("post object: %w", err)
	}

	var messageCreate messages.MessageCreate
	json.Unmarshal(bodyResponse, &messageCreate)

	return messageCreate.ID, nil
}

func PostFollowUpWithFile(token string, fileBytes []byte, filename string, interaction *builders.InteractionCallbackBuilder) error {

	interaction.
		AddAttachment(
			builders.CreateAttachment(0, "video", filename))

	writer, bodyBytes, err := http.CreateFormFileWithMessage(interaction, filename, fileBytes)
	if err != nil {
		return fmt.Errorf("create form file with message: %w", err)
	}

	_, _, err = client.DoRequest(getWebHookEndpointForToken(token), "POST", writer.FormDataContentType(), bodyBytes)
	if err != nil {
		return fmt.Errorf("post: %w", err)
	}

	return nil
}

func PostFollowUp(id string, token string, interactionCallbackPayload *messages.InteractionCallback) error {

	_, _, err := client.DoPostObject(getWebHookEndpointForToken(token), interactionCallbackPayload)
	if err != nil {
		return fmt.Errorf("post object: %w", err)
	}

	return nil
}

func GetOriginalInteraction(appId string, token string, messageId string) (*messages.InteractionCallback, error) {

	_, bodyResponse, err := client.DoGet(getOriginalMessageEndpointForToken(token))
	if err != nil {
		return nil, fmt.Errorf("delete: %w", err)
	}

	var interactionCallbackPayload messages.InteractionCallback
	err = json.Unmarshal(bodyResponse, &interactionCallbackPayload)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %s", err)
	}

	return &interactionCallbackPayload, nil
}

func EditOriginalInteraction(appId string, token string, interactionCallbackPayload *messages.Data) error {

	_, _, err := client.DoPatchObject(getOriginalMessageEndpointForToken(token), interactionCallbackPayload)
	if err != nil {
		return fmt.Errorf("patch object: %w", err)
	}

	return nil
}

func DeleteOriginalInteraction(token string) error {

	_, _, err := client.DoDelete(getOriginalMessageEndpointForToken(token))
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}
