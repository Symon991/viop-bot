package http

import (
	"bot/utils/builders"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
)

type Client struct {
	Headers map[string]string
}

type ClientInterface interface {
	DoPatchObject(url string, object interface{}) (int, []byte, error)
	DoPostObject(url string, object interface{}) (int, []byte, error)
	DoGet(url string) (int, []byte, error)
	DoDelete(url string) (int, []byte, error)
	DoRequest(url string, method string, contentType string, body []byte) (int, []byte, error)
}

func (client *Client) createRequest(url string, method string, contentType string, body []byte) (*http.Request, error) {

	request, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}

	for header, value := range client.Headers {
		request.Header.Set(header, value)
	}

	return request, nil
}

func (client *Client) DoRequest(url string, method string, contentType string, body []byte) (int, []byte, error) {

	request, err := client.createRequest(url, method, contentType, body)
	if err != nil {
		return http.StatusInternalServerError, nil, fmt.Errorf("request create: %w", err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return http.StatusInternalServerError, nil, fmt.Errorf("post message: %s", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return http.StatusInternalServerError, nil, fmt.Errorf("read response: %w", err)
	}

	log.Printf("%s", string(responseBody))

	return response.StatusCode, responseBody, nil
}

func (client *Client) DoStream(url string, method string, contentType string, body []byte) (*http.Response, error) {

	request, err := client.createRequest(url, method, contentType, body)
	if err != nil {
		return nil, fmt.Errorf("request create: %w", err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("post message: %s", err)
	}

	return response, nil
}

func (client *Client) DoPostObject(url string, object interface{}) (int, []byte, error) {

	return client.DoRequestObject(url, "POST", object)
}

func (client *Client) DoPatchObject(url string, object interface{}) (int, []byte, error) {

	return client.DoRequestObject(url, "PATCH", object)
}

func (client *Client) DoDelete(url string) (int, []byte, error) {

	return client.DoRequest(url, "DELETE", "", nil)
}

func (client *Client) DoGet(url string) (int, []byte, error) {

	return client.DoRequest(url, "GET", "", nil)
}

func (client *Client) DoRequestObject(url string, method string, object interface{}) (int, []byte, error) {

	body, err := json.Marshal(object)
	if err != nil {
		return http.StatusInternalServerError, nil, fmt.Errorf("marshal object: %w", err)
	}

	return client.DoRequest(url, method, "application/json", body)
}

func CreateFormFileWithMessage(interactionResponse *builders.InteractionCallbackBuilder, filename string, fileBytes []byte) (*multipart.Writer, []byte, error) {

	interactionResponseBytes, err := json.Marshal(interactionResponse.Get().Data)
	if err != nil {
		return nil, nil, fmt.Errorf("marshal interaction response: %w", err)
	}

	body := bytes.Buffer{}
	writer := multipart.NewWriter(&body)

	mime := textproto.MIMEHeader{}
	mime.Set("Content-Disposition", "form-data; name=\"payload_json\"")
	mime.Set("Content-Type", "application/json")

	contentWriter, err := writer.CreatePart(mime)
	if err != nil {
		return nil, nil, fmt.Errorf("create form: %w", err)
	}

	_, err = contentWriter.Write(interactionResponseBytes)
	if err != nil {
		return nil, nil, fmt.Errorf("write form part: %w", err)
	}

	fileWriter, err := writer.CreateFormFile("files[0]", filename)
	if err != nil {
		return nil, nil, fmt.Errorf("create form file: %w", err)
	}

	_, err = fileWriter.Write(fileBytes)
	if err != nil {
		return nil, nil, fmt.Errorf("write form file: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, nil, fmt.Errorf("close writer: %w", err)
	}

	return writer, body.Bytes(), nil
}
