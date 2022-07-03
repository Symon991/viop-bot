package main

import (
	"bot/api"
)

func main() {

	conn, interval := api.Connect()
	defer conn.Close()

	api.Heartbeat(interval, conn)
	api.Identify(conn)
	api.Listen(conn)
}
