package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ryanbaskara/learning-go/config"
)

func main() {
	grpcServer, err := config.NewGRPCServer()
	if err != nil {
		log.Fatal(err.Error())
	}

	listen, err := net.Listen("tcp", grpcServer.RPCHost)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func(g *config.GRPCServer, lis net.Listener) {
		log.Printf("grpc Server is available at %s\n", g.RPCHost)
		if err := g.Server.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}(grpcServer, listen)

	<-sigChan

	log.Println("shutting down the grpc Server...")

	if err := grpcServer.Stop(); err != nil {
		log.Println("error in shutting down grpc server: ", err)
	} else {
		log.Println("grpc Server gracefully stopped")
	}
}
