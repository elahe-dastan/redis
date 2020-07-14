package main

import (
	//import the redigo/redis package

	"fmt"
	"log"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

// Define a custom struct to hold album data. Notice the struct tags?
// These indicate to redigo how to assign the data from the reply into
// the struct.
type Album struct {
	Title string  `redis:"title"`
	Artist string  `redis:"artist"`
	Price float64  `redis:"price"`
	Likes int      `redis:"likes"`
}

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

	//title, err := redis.String(conn.Do("HGET", "album2", "title"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//artist, err := redis.String(conn.Do("HGET", "album2", "artist"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// And the price as a float64...
	//price, err := redis.Float64(conn.Do("HGET", "album2", "price"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// And the number of likes as an integer.
	//likes, err := redis.Int(conn.Do("HGET", "album2", "likes"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("%s by %s: £%.2f [%d likes]\n", title, artist, price, likes)

	// Fetch all album fields with the HGETALL command. Because HGETALL
	// returns an array reply,and because the underlying data structure
	// in Redis is a hash, it makes sense to use the Map() helper
	// function to convert the reply to a map[string]string.
	//reply, err := redis.StringMap(conn.Do("HGETALL", "album2"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//album, err := populateAlbum(reply)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("%+v", album)

	// Fetch all album fields with the HGETALL command. Wrapping this
	// in the redis.Values() function transforms the response into type
	// []interface{}, which is the format we need to pass to
	// redis.ScanStruct() in the next step.
	values, err := redis.Values(conn.Do("HGETALL", "album2"))
	if err != nil {
		log.Fatal(err)
	}

	var album Album
	err = redis.ScanStruct(values, &album)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", album)
}

func populateAlbum(reply map[string]string) (*Album, error) {
	var err error
	album := new(Album)
	album.Title = reply["title"]
	album.Artist = reply["artist"]
	// We need to use the strconv package to convert the 'price' value
	// from a string to a float64 before assigning it.
	album.Price, err = strconv.ParseFloat(reply["price"], 64)
	if err != nil {
		return nil, err
	}

	album.Likes, err = strconv.Atoi(reply["likes"])
	if err != nil {
		return nil, err
	}
	return album, nil
}