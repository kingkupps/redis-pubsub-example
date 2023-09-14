package main

import (
	"log"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

type Message struct {
	Receiver string `json:"receiver"`
	Msg      string `json:"message"`
}

func main() {
	redisClient := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_ENDPOINT")})
	http.Handle("/send", SendHandler{redisClient})
	http.Handle("/listen", ListenHandler{redisClient})
	if err := http.ListenAndServe(":1932", nil); err != nil {
		log.Fatalf("failed to start server: %s\n", err)
	}
}
