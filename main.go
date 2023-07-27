package main

import (
	"context"
	"fmt"
	handler "github.com/krylphi/slice-tryout/handler"
	repo "github.com/krylphi/slice-tryout/repo"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {

	serverAddr := os.Getenv("ADDR")
	if serverAddr == "" {
		log.Fatal("ADDR value is missing")
	}
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	store := repo.NewStore()

	srv := handler.NewHandler(store, serverAddr, serverPort)

	// Create a channel to listen for SIGINT and SIGTERM signals.
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, os.Kill)

	// Start the server on the local IP address and port 8080.
	go func() {
		if err := srv.Start(); err != nil {
			fmt.Println(err)
		}
	}()

	// Listen for the signal and gracefully shutdown the server.
	<-done
	fmt.Println("Received signal, shutting down gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Server exiting...")

}
