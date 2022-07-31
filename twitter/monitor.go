package twitter

import (
	"bot/discord"
	"bot/discord/messages"
	"fmt"
	"strings"
)

func Monitor(errorChan chan error) {

	channel := make(chan StreamMessage, 1)
	go Stream(channel, errorChan)

	for {
		tweet := (<-channel)
		var matchingRules []string
		for _, matchingRule := range tweet.MatchingRules {
			matchingRules = append(matchingRules, matchingRule.Tag)
		}
		err := discord.PostChannelMessage(messages.ChannelMessage{
			Content: fmt.Sprintf("[%s] https://twitter.com/user/status/%s", strings.Join(matchingRules, ", "), tweet.Data.ID),
		})
		if err != nil {
			errorChan <- fmt.Errorf("post tweet on channel: %w", err)
			return
		}
	}
}
