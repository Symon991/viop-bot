package discord

import (
	"bot/discord/messages"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func PostInteractionCallback(id string, token string, interactionCallbackPayload *messages.InteractionCallback) error {

	callbackPayload, err := json.Marshal(interactionCallbackPayload)
	if err != nil {
		return fmt.Errorf("marshal callback message: %s", err)
	}

	callback := fmt.Sprintf(discordCallbackTemplateUrl, id, token)
	log.Printf("%s\n\n", callback)
	log.Printf("%s\n\n", callbackPayload)
	response, err := http.Post(callback, "application/json", bytes.NewBuffer(callbackPayload))
	if err != nil {
		return fmt.Errorf("post callback message: %s", err)
	}

	body, _ := io.ReadAll(response.Body)
	log.Printf("%s\n\n", string(body))

	return nil
}

func PostFollowUp(id string, token string, interactionCallbackPayload *messages.InteractionCallback) error {

	callbackPayload, err := json.Marshal(interactionCallbackPayload)
	if err != nil {
		return fmt.Errorf("marshal follow up message: %s", err)
	}

	callback := fmt.Sprintf(discordFollowUpTemplateUrl, id, token)
	log.Printf("%s\n\n", callback)
	log.Printf("%s\n\n", callbackPayload)
	response, err := http.Post(callback, "application/json", bytes.NewBuffer(callbackPayload))
	if err != nil {
		return fmt.Errorf("post follow up message: %s", err)
	}

	body, _ := io.ReadAll(response.Body)
	fmt.Printf("%s\n\n", string(body))

	return nil
}

func GetOriginalInteraction(appId string, token string, messageId string) (*messages.InteractionCallback, error) {

	getCallback := fmt.Sprintf(discordGetCallbackTemplateUrl, appId, token)
	response, err := http.Get(getCallback)
	log.Printf("%s\n\n", getCallback)
	if err != nil {
		return nil, fmt.Errorf("get original interaction: %s", err)
	}
	body, _ := io.ReadAll(response.Body)
	log.Printf("debug getOriginal Interaction %s\n\n", body)

	var interactionCallbackPayload messages.InteractionCallback
	err = json.Unmarshal(body, &interactionCallbackPayload)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %s", err)
	}

	return &interactionCallbackPayload, nil
}

func EditOriginalInteraction(appId string, token string, interactionCallbackPayload *messages.Data) error {

	callbackPayload, err := json.Marshal(interactionCallbackPayload)
	if err != nil {
		return fmt.Errorf("marshal callback message: %s", err)
	}

	callback := fmt.Sprintf(discordEditCallbackTemplateUrl, appId, token)
	log.Printf("%s\n\n", callback)
	log.Printf("%s\n\n", callbackPayload)

	request, _ := http.NewRequest("PATCH", callback, bytes.NewBuffer(callbackPayload))
	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("post callback message: %s", err)
	}

	body, _ := io.ReadAll(response.Body)
	log.Printf("%s", string(body))

	return nil
}
