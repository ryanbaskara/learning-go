package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ryanbaskara/learning-go/config"
)

func main() {
	server, err := config.NewServer()
	if err != nil {
		log.Fatalln("Error when configure server : ", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func(s *http.Server) {
		log.Printf("Server is available at %s\n", s.Addr)
		if serr := s.ListenAndServe(); serr != http.ErrServerClosed {
			log.Fatalln(serr)
		}
	}(server.HttpServer)

	<-sigChan

	log.Println("Shutting down server...")
	if err := server.HttpServer.Shutdown(context.Background()); err != nil {
		log.Fatalln("Something wrong when stopping server : ", err)
		return
	}

	log.Println("Server gracefully stopped")
}
