package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type SendHandler struct {
	redisClient *redis.Client
}

func (s SendHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var msg Message
	if err := decoder.Decode(&msg); err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "failed to parse message body: %s", err)
		return
	}
	// You shouldn't put large payloads as messages in topics. If you have large
	// payloads (>100KB), you should publish an identifier to the payload, then
	// have the receiver lookup the payload by the identifier.
	if _, err := s.redisClient.Publish(r.Context(), msg.Receiver, msg.Msg).Result(); err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "failed to publish message: %s", err)
		return
	}
	fmt.Fprint(w, "sent message :)\n")
}
