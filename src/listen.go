package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

type ListenHandler struct {
	redisClient *redis.Client
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

func (l ListenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	receiver := r.URL.Query().Get("receiver")
	if receiver == "" {
		w.WriteHeader(400)
		fmt.Fprint(w, "missing receiver parameter")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "failed to upgrade connection: %s", err)
		return
	}

	sub := l.redisClient.Subscribe(r.Context(), receiver)
	defer sub.Close()
	for {
		next := <-sub.Channel()
		if err := conn.WriteMessage(websocket.TextMessage, []byte(next.Payload)); err != nil {
			log.Printf("failed to write back to client: %s", err)
			conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}
	}
}
