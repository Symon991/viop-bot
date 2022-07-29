package main

import (
	"bot/cache"
	"bot/commands"
	"bot/discord"
	"bot/discord/messages"
	"bot/twitter"
	"fmt"
	"net/url"
	"os"
)

func main() {

	fmt.Println("START")

	conn, interval := discord.Connect()
	defer conn.Close()

	redisCloudUrlEnv := os.Getenv("REDISCLOUD_URL")
	redisCloudUrl, err := url.Parse(redisCloudUrlEnv)
	if err != nil {
		panic(fmt.Errorf("paring redis cloud url: %w", err))
	}

	cache.Connect(redisCloudUrl)

	discord.Heartbeat(interval, conn)
	discord.Identify(conn, os.Getenv("DISCORD_APPLICATION_ID"))
	discord.Listen(conn, commands.HandleInteraction)

	channel := make(chan twitter.StreamMessage, 1)
	go twitter.Stream(channel)

	for {
		discord.PostChannelMessage(messages.ChannelMessage{
			Content: (<-channel).Data.Text,
		})
	}
}
