package main

import (
	"golang.org/x/net/context"

	example "github.com/hnakamur/grpc-gateway-example"
)

type Server struct{}

func (s *Server) Echo(ctx context.Context, msg *example.StringMessage) (*example.StringMessage, error) {
	return msg, nil
}
