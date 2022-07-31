package twitter

import (
	"bot/discord"
	"bot/discord/messages"
	"fmt"
	"strings"
)

func Monitor(errorChan chan error) {

	channel := make(chan StreamMessage, 1)

	streamErrorChan := make(chan error)
	go Stream(channel, streamErrorChan)

	for {
		select {
		case err := <-streamErrorChan:
			{
				errorChan <- fmt.Errorf("twitter stream: %w", err)
			}
		default:
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
}
