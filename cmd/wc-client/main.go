package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/context"

	"github.com/ljsc/wcsvc"
	"google.golang.org/grpc"
)

func main() {
	text := strings.Join(os.Args[1:len(os.Args)], " ")
	conn, err := grpc.Dial("localhost:4242", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	wc := wordcount.NewWordCountClient(conn)
	stats, err := wc.GetStats(context.Background(), &wordcount.Message{Text: text})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Stats: %+v\n", stats)
}
