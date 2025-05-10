package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/cmd/order/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		sig := <-sigChan
		log.Printf("⚠️ Received signal: %v. Shutting down...\n", sig)
		cancel()
	}()

	server, err := app.NewServer()
	if err != nil {
		log.Fatalf("❌ Failed to initialize server: %v", err)
	}

	<-ctx.Done()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("❌ Failed to shutdown server: %v", err)
	}

	err = server.MongoDB.Close()
	if err != nil {
		log.Fatalf("❌ Failed to close database: %v", err)
	}

	log.Println("✅ Server shutdown gracefully")
}
