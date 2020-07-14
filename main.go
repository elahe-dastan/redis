package main

import (
	//import the redigo/redis package

	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)
func main() {
	//Establish a connection to yhe redis server listening on the default port 6379
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	//_, err = conn.Do("HMSET", "album2", "title", "Electric ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
	//if err != nil {
	//	log.Fatal(err)
	//}

	title, err := redis.String(conn.Do("HGET", "album2", "title"))
	if err != nil {
		log.Fatal(err)
	}

	artist, err := redis.String(conn.Do("HGET", "album2", "artist"))
	if err != nil {
		log.Fatal(err)
	}

	// And the price as a float64...
	price, err := redis.Float64(conn.Do("HGET", "album2", "price"))
	if err != nil {
		log.Fatal(err)
	}

	// And the number of likes as an integer.
	likes, err := redis.Int(conn.Do("HGET", "album2", "likes"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s by %s: £%.2f [%d likes]\n", title, artist, price, likes)
}
