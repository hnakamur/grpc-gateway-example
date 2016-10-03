package main

import (
	"log"
	"net"

	example "github.com/hnakamur/grpc-gateway-example"

	"google.golang.org/grpc"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	example.RegisterYourServiceServer(s, new(Server))

	s.Serve(l)
	return nil
}
