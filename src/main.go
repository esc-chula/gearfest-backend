package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/esc-chula/gearfest-backend/src/server"
)

// @title Gearfest API
// @version 1.0
// @description API for the GearFestival website
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and Google ID token.
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
