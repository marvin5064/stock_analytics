package main

import (
	"context"
	"fmt"
	"log"
	"net"

	stock "github.com/marvin5064/stock-analytics/protobuf/stock"
	"google.golang.org/grpc/reflection"
)

func RunGrpcServer(srv *Server, host string, port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	stock.RegisterStockManagerServer(srv.grpcSrv, srv)
	// Register reflection service on gRPC server.
	reflection.Register(srv.grpcSrv)
	if err := srv.grpcSrv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) GetStockPrices(
	ctx context.Context,
	request *stock.StockPriceRequest) (*stock.StockPriceResponse, error) {
	return &stock.StockPriceResponse{}, nil

}
