package main

import (
	//import the redigo/redis package

	"log"

	"github.com/gomodule/redigo/redis"
)
func main() {
	//Establish a connection to yhe redis server listening on the default port 6379
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Do("HMSET", "album2", "title", "Electric ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
}
