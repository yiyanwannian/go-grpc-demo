package main

import (
	"context"
	"fmt"
	"github.com/albinism/grpc-test/protos/hello"
	"google.golang.org/grpc"
	"log"
	"sync"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)

	wg := &sync.WaitGroup{}

	for _, n := range []string{"tyler", "roscoe"} {
		wg.Add(1)
		go func(name string) {
			response, err := client.Hello(context.Background(), &hello.HelloRequest{Name: name})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(response)
			wg.Done()
		}(n)
	}

	wg.Wait()
}
