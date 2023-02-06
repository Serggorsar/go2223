package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"

	"microservices/grpc_stream/translit"
)

func main() {

	grcpConn, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

	tr := translit.NewTransliterationClient(grcpConn)

	ctx := context.Background()
	client, err := tr.EnRu(ctx)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		words := []string{"privet", "kak", "dela"}
		for _, w := range words {
			fmt.Println("in:", w)
			client.Send(&translit.Word{
				Word: w,
			})
			time.Sleep(time.Millisecond)
		}
		client.CloseSend()
		fmt.Println("\tsend done")
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			outWord, err := client.Recv()
			if err == io.EOF {
				fmt.Println("stream closed")
				return
			} else if err != nil {
				fmt.Println("error happed", err)
				return
			}
			fmt.Println("out:", outWord.Word)
		}
		fmt.Println("recv done")
	}(wg)

	wg.Wait()

}
