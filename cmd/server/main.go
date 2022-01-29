package main

import (
	api "getputer/pkg/api/versions/v1"
	"getputer/pkg/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	s := grpc.NewServer()
	srv := &server.GRPCServer{}
	api.RegisterGet_PutServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
