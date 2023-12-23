package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/esc-chula/gearfest-backend/src/server"
)

func main() {
	server := server.New()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, os.Interrupt, syscall.SIGTERM)
	defer cancel()
	err := server.Start(ctx)
	if err != nil {
		fmt.Println("Error on server, shutting down.")
	}
	fmt.Println("Exit.")
	os.Exit(0)
}
