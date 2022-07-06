package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func PostInteractionCallback(id string, token string, interactionCallbackPayload *InteractionCallbackPayload) error {

	callbackPayload, err := json.Marshal(interactionCallbackPayload)
	if err != nil {
		return fmt.Errorf("error marshaling callback: %s", err)
	}

	callback := fmt.Sprintf(discordCallbackTemplateUrl, id, token)
	fmt.Printf("%s\n\n", callback)
	fmt.Printf("%s\n\n", callbackPayload)
	response, err := http.Post(callback, "application/json", bytes.NewBuffer(callbackPayload))
	if err != nil {
		return fmt.Errorf("error during post to callback: %s", err)
	}

	body, _ := io.ReadAll(response.Body)
	fmt.Printf("%s\n\n", string(body))

	return nil
}

func GetOriginalInteraction(appId string, token string, messageId string) (*InteractionCallbackPayload, error) {

	getCallback := fmt.Sprintf(discordGetCallbackTemplateUrl, appId, token)
	response, err := http.Get(getCallback)
	fmt.Printf("%s\n\n", getCallback)
	if err != nil {
		return nil, fmt.Errorf("error getting original message: %s", err)
	}
	body, _ := io.ReadAll(response.Body)
	fmt.Printf("debug getOriginal Interaction %s\n\n", body)

	var interactionCallbackPayload InteractionCallbackPayload
	err = json.Unmarshal(body, &interactionCallbackPayload)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %s", err)
	}

	return &interactionCallbackPayload, nil
}

func EditOriginalInteraction(appId string, token string, messageId string, interactionCallbackPayload *InteractionCallbackPayload) error {

	callbackPayload, err := json.Marshal(interactionCallbackPayload)
	if err != nil {
		return fmt.Errorf("error marshaling callback: %s", err)
	}

	callback := fmt.Sprintf(discordEditCallbackTemplateUrl, appId, token)
	fmt.Printf("%s\n\n", callback)
	fmt.Printf("%s\n\n", callbackPayload)

	request, _ := http.NewRequest("PATCH", callback, bytes.NewBuffer(callbackPayload))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("error during post to callback: %s", err)
	}

	body, _ := io.ReadAll(response.Body)
	fmt.Printf("%s", string(body))

	return nil
}
