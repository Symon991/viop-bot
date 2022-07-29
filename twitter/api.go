package twitter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const tweetStreamFilterUrl = "https://api.twitter.com/2/tweets/search/stream/rules"
const tweetStreamUrl = "https://api.twitter.com/2/tweets/search/stream"

type AddRuleMessage struct {
	Add []Rule `json:"add"`
}

type DeleteRuleMessage struct {
	Delete struct {
		IDs []string `json:"ids"`
	} `json:"delete"`
}

type GetRulesMessage struct {
	Data []Rule `json:"data"`
}

type Rule struct {
	ID    string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
	Tag   string `json:"tag,omitempty"`
}

type StreamMessage struct {
	Data Data `json:"data"`
}

type Data struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func AddRule(value string, tag string) error {

	rule := Rule{
		Value: value,
		Tag:   tag,
	}

	var addRuleMessage AddRuleMessage
	addRuleMessage.Add = append(addRuleMessage.Add, rule)

	byte, err := json.Marshal(addRuleMessage)
	if err != nil {
		return fmt.Errorf("marshal addRuleMessage: %w", err)
	}
	fmt.Printf("%s", byte)

	request, err := http.NewRequest("POST", tweetStreamFilterUrl, bytes.NewBuffer(byte))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAAD%2FGfAEAAAAAcz3b%2FgLeLDFSM%2Fgvmd2MxHXqUD4%3DXkm65WicZo7jKVeKu3tBy5CSZmvbRdoHceX6WnJCFuJbKCl5Jf")

	response, err := (&http.Client{}).Do(request)
	if err != nil {
		return fmt.Errorf("post add rule: %w", err)
	}
	byte, _ = io.ReadAll(response.Body)
	fmt.Printf("%s", byte)

	return nil
}

func GetRules() (*GetRulesMessage, error) {

	request, err := http.NewRequest("GET", tweetStreamFilterUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	request.Header.Set("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAAD%2FGfAEAAAAAcz3b%2FgLeLDFSM%2Fgvmd2MxHXqUD4%3DXkm65WicZo7jKVeKu3tBy5CSZmvbRdoHceX6WnJCFuJbKCl5Jf")

	response, err := (&http.Client{}).Do(request)
	if err != nil {
		return nil, fmt.Errorf("get rules: %w", err)
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response get rules: %w", err)
	}

	var getRulesMessage GetRulesMessage
	err = json.Unmarshal(bytes, &getRulesMessage)
	if err != nil {
		return nil, fmt.Errorf("unmarshal get rules: %w", err)
	}

	return &getRulesMessage, nil
}

func RemoveRule(id string) error {

	var deleteRuleMessage DeleteRuleMessage
	deleteRuleMessage.Delete.IDs = append(deleteRuleMessage.Delete.IDs, id)

	byte, err := json.Marshal(deleteRuleMessage)
	if err != nil {
		return fmt.Errorf("marshal deleteRuleMessage: %w", err)
	}
	fmt.Printf("%s", byte)

	request, err := http.NewRequest("POST", tweetStreamFilterUrl, bytes.NewBuffer(byte))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAAD%2FGfAEAAAAAcz3b%2FgLeLDFSM%2Fgvmd2MxHXqUD4%3DXkm65WicZo7jKVeKu3tBy5CSZmvbRdoHceX6WnJCFuJbKCl5Jf")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("post add rule: %w", err)
	}
	byte, _ = io.ReadAll(response.Body)
	fmt.Printf("%s", byte)

	return nil
}

func Stream(dataChannel chan StreamMessage) error {

	request, err := http.NewRequest("GET", tweetStreamUrl, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	request.Header.Set("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAAD%2FGfAEAAAAAcz3b%2FgLeLDFSM%2Fgvmd2MxHXqUD4%3DXkm65WicZo7jKVeKu3tBy5CSZmvbRdoHceX6WnJCFuJbKCl5Jf")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("get rules: %w", err)
	}

	var streamMessage StreamMessage
	for {
		err := json.NewDecoder(response.Body).Decode(&streamMessage)
		if err != nil {
			return fmt.Errorf("error stream: %w", err)
		}
		dataChannel <- streamMessage
	}
}
