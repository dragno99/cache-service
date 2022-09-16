package main

import (
	"time"

	"github.com/dragno99/cache-service/client"
	"github.com/dragno99/cache-service/server"
)

// var ctx = context.Background()

// func ExampleClient() {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})

// 	err := rdb.Set("key", "value", 0).Err()
// 	if err != nil {
// 		panic(err)
// 	}

// 	val, err := rdb.Get("key").Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("key", val)

// 	val2, err := rdb.Get("key2").Result()
// 	if err == redis.Nil {
// 		fmt.Println("key2 does not exist")
// 	} else if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("key2", val2)
// 	}
// 	// Output: key value
// 	// key2 does not exist
// }

func main() {

	// starting server in another go routine so that it wont block our code
	go server.StartServer()

	time.Sleep(time.Millisecond * 5000)

	// calling this method to test the client
	client.Test()

	<-make(chan struct{})
}
