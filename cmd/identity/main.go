package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/cmd/identity/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go listenForShutdownSignal(cancel)

	server, err := initializeServer()
	if err != nil {
		log.Fatalf("❌ Failed to initialize server: %v", err)
	}

	<-ctx.Done()

	shutdownServer(server)
	log.Println("✅ Server shutdown gracefully")
}

func listenForShutdownSignal(cancel context.CancelFunc) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	sig := <-sigChan
	log.Printf("⚠️ Received signal: %v. Shutting down...\n", sig)
	cancel()
}

func initializeServer() (*app.Server, error) {
	server, err := app.NewServer()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize server: %v", err)
	}
	return server, nil
}

func shutdownServer(server *app.Server) {
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("❌ Failed to shutdown server: %v", err)
	}

	// Close the MongoDB connection
	err := server.MongoDB.Close()
	if err != nil {
		log.Fatalf("❌ Failed to close database: %v", err)
	}
}
