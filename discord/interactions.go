package discord

import (
	"bot/discord/messages"
	"bot/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
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
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	log.Printf("%s\n\n", string(body))

	return nil
}

func PostInteractionFile(id string, token string, fileBytes []byte) error {

	interaction := utils.CreateInteractionCallback().AddAttachment(utils.CreateAttachment(0, "video", "video.webm"))

	//callback := fmt.Sprintf(discordCallbackTemplateUrl, id, token)

	payloadJsonBytes, err := json.Marshal(interaction)
	if err != nil {
		return fmt.Errorf("marshal payload: %w", err)
	}

	log.Println("post interaction file")

	body := bytes.Buffer{}
	writer := multipart.NewWriter(&body)

	log.Println("write content")
	contentWriter, err := writer.CreatePart("json_payload")
	if err != nil {
		return fmt.Errorf("create form: %w", err)
	}
	contentWriter.Write(payloadJsonBytes)

	log.Println("write file")
	fileWriter, err := writer.CreateFormFile("files[0]", "video.webm")
	if err != nil {
		return fmt.Errorf("create form file: %w", err)
	}
	_, err = fileWriter.Write(fileBytes)
	if err != nil {
		return fmt.Errorf("write file form: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return fmt.Errorf("close writer: %w", err)
	}

	log.Println("writer closed")

	request, err := http.NewRequest("POST", discordPostChannelBotInfoUrl, bytes.NewReader(body.Bytes()))
	if err != nil {
		return fmt.Errorf("request create: %w", err)
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("DISCORD_APPLICATION_ID")))

	log.Println("request")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("request create: %w", err)
	}

	defer response.Body.Close()

	responseBody, _ := io.ReadAll(response.Body)
	log.Printf("%s", string(responseBody))

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
	defer response.Body.Close()

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
	defer response.Body.Close()

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
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	log.Printf("%s", string(body))

	return nil
}
