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
)

func main() {

	log.Println("START")

	redisCloudUrlEnv := os.Getenv("REDISCLOUD_URL")
	redisCloudUrl, err := url.Parse(redisCloudUrlEnv)
	if err != nil {
		log.Panicln(fmt.Errorf("paring redis cloud url: %w", err))
	}
	cache.Connect(redisCloudUrl)

	discordErrorChan := make(chan error)
	twitterErrorChan := make(chan error)

	conn, interval := discord.Connect()
	discord.Heartbeat(interval, conn, discordErrorChan)
	discord.Identify(conn, os.Getenv("DISCORD_APPLICATION_ID"))
	go discord.Listen(conn, commands.HandleInteraction, discordErrorChan)

	go twitter.Monitor(twitterErrorChan)

	for {
		select {
		case <-discordErrorChan:
			log.Println(err)

			conn.Close()

			conn, interval := discord.Connect()
			discord.Heartbeat(interval, conn, discordErrorChan)
			discord.Identify(conn, os.Getenv("DISCORD_APPLICATION_ID"))
			go discord.Listen(conn, commands.HandleInteraction, discordErrorChan)

		case <-twitterErrorChan:
			go twitter.Monitor(twitterErrorChan)
		}
	}
}
