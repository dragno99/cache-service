package client

import (
	"context"
	"fmt"
	"log"

	proto "github.com/dragno99/cache-service/proto"

	"google.golang.org/grpc"
)

func Test() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := proto.NewAddServiceClient(conn)

	// testing Set method
	req1 := &proto.KeyVal{
		Key:   "suru",
		Value: []byte("here"),
	}

	_, err = client.Set(context.Background(), req1)
	if err != nil {
		fmt.Println("Response recieved by Set method:", err.Error())
	}

	// testing Get method
	req2 := &proto.Key{Key: "suruIshere"}
	res2, err := client.Get(context.Background(), req2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Response recieved by Get method:", string(res2.GetValue()))

}
