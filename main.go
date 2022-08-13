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

	"github.com/gorilla/websocket"
)

func main() {

	discordErrorChan := make(chan error)
	twitterErrorChan := make(chan error)

	startRedis()
	conn := startDiscord(discordErrorChan)
	defer conn.Close()
	startTwitter(twitterErrorChan)

	for {
		select {
		case err := <-discordErrorChan:
			log.Print(err)
			//log.Print("connection error detected, reconnect in 5 seconds")
			//conn.Close()

			//time.Sleep(time.Second * 30)
			//conn = startDiscord(discordErrorChan)

		case err := <-twitterErrorChan:
			log.Print(err)
			//log.Print("twitter monitor error detected, restart in 5 seconds")
			//time.Sleep(time.Second * 30)
			//startTwitter(twitterErrorChan)
		}
	}
}

func startDiscord(errorChan chan error) *websocket.Conn {

	conn, interval, err := discord.Connect()
	if err != nil {
		log.Panic(err)
	}
	discord.Heartbeat(interval, conn, errorChan)
	discord.Identify(conn, os.Getenv("DISCORD_BEARER_TOKEN"))
	go discord.Listen(conn, commands.HandleInteraction, errorChan)
	return conn
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
