package main

import (
	"time"

	client "github.com/dragno99/cache-service/client"
	server "github.com/dragno99/cache-service/server"
)

func main() {

	// starting server in another go routine so that it wont block our code
	go server.StartServer()

	time.Sleep(time.Millisecond * 5000)

	// calling this method to test the client
	client.Test()

	<-make(chan struct{})
}
