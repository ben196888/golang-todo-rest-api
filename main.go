package main

import (
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func main() {
	router := NewRouter()
	log.Println("API server listen on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
