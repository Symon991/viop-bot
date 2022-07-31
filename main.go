package main

import (
	"bot/cache"
	"bot/commands"
	"bot/discord"
	"bot/twitter"
	"fmt"
	"log"
	"net/url"
	"os"

	"golang.org/x/net/websocket"
)

func main() {

	discordErrorChan := make(chan error)
	twitterErrorChan := make(chan error)

	conn := startDiscord(discordErrorChan)
	defer conn.Close()
	startTwitter(twitterErrorChan)

	for {
		select {
		case err := <-discordErrorChan:
			log.Print(err)
			conn.Close()
			conn = startDiscord(discordErrorChan)

		case err := <-twitterErrorChan:
			log.Print(err)
			startTwitter(twitterErrorChan)
		}
	}
}

func startDiscord(errorChan chan error) *websocket.Conn {

	conn, interval := discord.Connect()
	discord.Heartbeat(interval, conn, errorChan)
	discord.Identify(conn, os.Getenv("DISCORD_APPLICATION_ID"))
	go discord.Listen(conn, commands.HandleInteraction, errorChan)
}

func startTwitter(errorChan chan error) {

	go twitter.Monitor(errorChan)
}

func startRedis() {

	redisCloudUrlEnv := os.Getenv("REDISCLOUD_URL")
	redisCloudUrl, err := url.Parse(redisCloudUrlEnv)
	if err != nil {
		log.Panic(fmt.Errorf("paring redis cloud url: %w", err))
	}
	cache.Connect(redisCloudUrl)
}
