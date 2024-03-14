package main

import (
	"chat-app/websocket"
	"log"
	"net/http"
)

func serverWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()

}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()
	log.Println("Started Server")

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWS(pool, w, r)
	})

}

func main() {
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}
