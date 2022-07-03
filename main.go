package main

import (
	"bot/api"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	go func() {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()

	conn, interval := api.Connect()
	defer conn.Close()

	api.Heartbeat(interval, conn)
	api.Identify(conn)
	api.Listen(conn)
}
