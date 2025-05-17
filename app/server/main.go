package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ryanbaskara/learning-go/config"
)

func main() {
	serverConfig := config.NewServerConfig()

	// log.Fatal(http.ListenAndServe(serverConfig.Host, serverConfig.Router))

	s := &http.Server{
		Addr:         serverConfig.Host,
		Handler:      serverConfig.Router,
		ReadTimeout:  310 * time.Second,
		WriteTimeout: 310 * time.Second,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func(s *http.Server) {
		log.Printf("Server is available at %s\n", s.Addr)
		if serr := s.ListenAndServe(); serr != http.ErrServerClosed {
			log.Fatalln(serr)
		}
	}(s)

	<-sigChan

	log.Println("Shutting down server...")
	if err := s.Shutdown(context.Background()); err != nil {
		log.Fatalln("Something wrong when stopping server : ", err)
		return
	}

	log.Println("Server gracefully stopped")
}
