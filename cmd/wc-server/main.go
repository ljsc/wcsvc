package main

import (
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/net/context"

	"github.com/ljsc/wcsvc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func init() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	grpclog.SetLogger(logger)
}

func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	defer func() {
		grpclog.Printf("Finished request %s in %v\n", info.FullMethod, time.Since(start))
	}()
	return handler(ctx, req)
}

func main() {
	sock, err := net.Listen("tcp", ":4242")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor))
	wordcount.RegisterWordCountServer(grpcServer, &wordcount.Server{})
	grpcServer.Serve(sock)
}
