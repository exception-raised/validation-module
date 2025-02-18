package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/exception-raised/src/server"
	"github.com/exception-raised/src/service"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	validationService := service.NewValidationService()

	s := server.NewServer(
		server.WithValidationService(validationService),
	)

	go func() {
		if err := s.ListenAndServe(":8080"); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")
	if err := s.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server shutdown failed: %+v", err)
	}
	log.Println("Server gracefully stopped")
}
