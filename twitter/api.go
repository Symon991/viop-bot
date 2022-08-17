package twitter

import (
	"bot/utils/http"
	"encoding/json"
	"fmt"
	"os"
)

const tweetStreamFilterUrl = "https://api.twitter.com/2/tweets/search/stream/rules"
const tweetStreamUrl = "https://api.twitter.com/2/tweets/search/stream"

var client *http.Client

func Init() {

	client = getClient()
}

func getClient() *http.Client {

	headers := make(map[string]string)

	headers["Authorization"] = fmt.Sprintf("Bearer %s", os.Getenv("TWITTER_BEARER_TOKEN"))

	return &http.Client{
		Headers: headers,
	}
}

func AddRule(value string, tag string) error {

	rule := Rule{
		Value: value,
		Tag:   tag,
	}

	var addRuleMessage AddRuleMessage
	addRuleMessage.Add = append(addRuleMessage.Add, rule)

	_, _, err := client.DoPostObject(tweetStreamFilterUrl, addRuleMessage)
	if err != nil {
		return fmt.Errorf("post object: %w", err)
	}

	return nil
}

func GetRules() (*GetRulesMessage, error) {

	_, bodyResponse, err := client.DoGet(tweetStreamFilterUrl)
	if err != nil {
		return nil, fmt.Errorf("post object: %w", err)
	}

	var getRulesMessage GetRulesMessage
	err = json.Unmarshal(bodyResponse, &getRulesMessage)
	if err != nil {
		return nil, fmt.Errorf("unmarshal get rules: %w", err)
	}

	return &getRulesMessage, nil
}

func RemoveRule(id string) error {

	var deleteRuleMessage DeleteRuleMessage
	deleteRuleMessage.Delete.IDs = append(deleteRuleMessage.Delete.IDs, id)

	_, _, err := client.DoPostObject(tweetStreamFilterUrl, deleteRuleMessage)
	if err != nil {
		return fmt.Errorf("post object: %w", err)
	}

	return nil
}

func Stream(dataChannel chan StreamMessage, errorChan chan error) {

	response, err := client.DoStream(tweetStreamUrl, "GET", "application/json", nil)
	if err != nil {
		errorChan <- fmt.Errorf("get rules: %w", err)
		return
	}

	var streamMessage StreamMessage
	for {
		err := json.NewDecoder(response.Body).Decode(&streamMessage)
		if err != nil {
			errorChan <- fmt.Errorf("error stream: %w", err)
			break
		}
		dataChannel <- streamMessage
	}
}
