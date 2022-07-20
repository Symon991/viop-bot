package main

import (
	"bot/cache"
	"bot/commands"
	"bot/discord"
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
}
