package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	pb "github.com/hikata101/health_data_service/gen/github.com/hikata101/health_data_service/v1"
	"github.com/hikata101/health_data_service/logger"
	"google.golang.org/grpc"

	handler "github.com/hikata101/health_data_service/interface"
)

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	logger.NewLogger()
	logger.Logger.Info("gRPC server is starting...")

	s := grpc.NewServer()
	server := handler.NewhealthDataServer()
	pb.RegisterDatasetServer(s, &server)

	go func() {
		logger.Logger.Info(fmt.Sprintf("start gRPC server port: %v", port))
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
