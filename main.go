package main

import (
	"bot/cache"
	"bot/commands"
	"bot/discord"
)

func main() {

	conn, interval := discord.Connect()
	defer conn.Close()

	cache.Connect()

	discord.Heartbeat(interval, conn)
	discord.Identify(conn)
	discord.Listen(conn, commands.HandleInteraction)
}
