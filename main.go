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
)

func main() {

	log.Println("START")

	conn, interval := discord.Connect()
	defer conn.Close()

	redisCloudUrlEnv := os.Getenv("REDISCLOUD_URL")
	redisCloudUrl, err := url.Parse(redisCloudUrlEnv)
	if err != nil {
		log.Panicln(fmt.Errorf("paring redis cloud url: %w", err))
	}

	cache.Connect(redisCloudUrl)

	discord.Heartbeat(interval, conn)
	discord.Identify(conn, os.Getenv("DISCORD_APPLICATION_ID"))
	go discord.Listen(conn, commands.HandleInteraction)

	channel := make(chan twitter.StreamMessage, 1)
	go twitter.Stream(channel)

	for {
		tweet := (<-channel)
		discord.PostChannelMessage(messages.ChannelMessage{
			Content: fmt.Sprintf("https://twitter.com/user/status/%s", tweet.Data.ID),
		})
	}
}
