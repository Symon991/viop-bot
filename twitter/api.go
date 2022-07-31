package twitter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	Data          Data           `json:"data"`
	MatchingRules []MatchingRule `json:"matching_rules"`
}

type MatchingRule struct {
	ID  string `json:"id"`
	Tag string `json:"tag"`
}

type Data struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func doRequest(request *http.Request) (*http.Response, error) {

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("TWITTER_BEARER_TOKEN")))

	return (&http.Client{}).Do(request)
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
	log.Printf("%s", byte)

	request, err := http.NewRequest("POST", tweetStreamFilterUrl, bytes.NewBuffer(byte))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	response, err := doRequest(request)
	if err != nil {
		return fmt.Errorf("post add rule: %w", err)
	}
	defer response.Body.Close()

	byte, _ = io.ReadAll(response.Body)
	log.Printf("%s", byte)

	return nil
}

func GetRules() (*GetRulesMessage, error) {

	request, err := http.NewRequest("GET", tweetStreamFilterUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	response, err := doRequest(request)
	if err != nil {
		return nil, fmt.Errorf("get rules: %w", err)
	}
	defer response.Body.Close()

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
	log.Printf("%s", byte)

	request, err := http.NewRequest("POST", tweetStreamFilterUrl, bytes.NewBuffer(byte))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	response, err := doRequest(request)
	if err != nil {
		return fmt.Errorf("post add rule: %w", err)
	}
	defer response.Body.Close()

	byte, _ = io.ReadAll(response.Body)
	log.Printf("%s", byte)

	return nil
}

func Stream(dataChannel chan StreamMessage, errorChan chan error) {

	request, err := http.NewRequest("GET", tweetStreamUrl, nil)
	if err != nil {
		errorChan <- fmt.Errorf("create request: %w", err)
		return
	}

	response, err := doRequest(request)
	if err != nil {
		errorChan <- fmt.Errorf("get rules: %w", err)
		return
	}

	var streamMessage StreamMessage
	for {
		err := json.NewDecoder(response.Body).Decode(&streamMessage)
		if err != nil {
			errorChan <- fmt.Errorf("error stream: %w", err)
			response.Body.Close()
			return
		}
		dataChannel <- streamMessage
	}
}
