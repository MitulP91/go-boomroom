package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	// "github.com/lib/pq"
)

// Define the address and port num for the redis instance
const (
	REDIS_ADDRESS = "127.0.0.1:6379"
)

// Main func. 
func main()	{
	fmt.Println("Chillin'. Like a master villain.")

	conn, err := redis.Dial("tcp", ":6379")
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	conn.Do("SET", "message1", "Hello, world!")

	world, err := redis.String(conn.Do("GET", "message1"))
	
	if err != nil {
		fmt.Println("key not found")
	}

	fmt.Println(world)
}