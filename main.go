package main

import (
	"bot/cache"
	"bot/commands"
	"bot/discord"
	"bot/discord/messages"
	"bot/twitter"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {

	log.Println("START")

	redisCloudUrlEnv := os.Getenv("REDISCLOUD_URL")
	redisCloudUrl, err := url.Parse(redisCloudUrlEnv)
	if err != nil {
		log.Panicln(fmt.Errorf("paring redis cloud url: %w", err))
	}
	cache.Connect(redisCloudUrl)

	conn, interval := discord.Connect()
	defer conn.Close()

	discord.Heartbeat(interval, conn)
	discord.Identify(conn, os.Getenv("DISCORD_APPLICATION_ID"))
	go discord.Listen(conn, commands.HandleInteraction)

	channel := make(chan twitter.StreamMessage, 1)
	go twitter.Stream(channel)

	for {
		tweet := (<-channel)
		var matchingRules []string
		for _, matchingRule := range tweet.MatchingRules {
			matchingRules = append(matchingRules, matchingRule.Tag)
		}
		discord.PostChannelMessage(messages.ChannelMessage{
			Content: fmt.Sprintf("[%s] https://twitter.com/user/status/%s", strings.Join(matchingRules, ", "), tweet.Data.ID),
		})
	}
}
