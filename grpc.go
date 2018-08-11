package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func (s *Server) setUpGrpcServer(host string, port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	stock.RegisterStockFetchManagerServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
