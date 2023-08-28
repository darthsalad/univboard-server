package main

import (
	"context"
	"time"
	"os"
	"os/signal"
	"syscall"

	"github.com/darthsalad/univboard/internal/logger"
	"github.com/darthsalad/univboard/pkg/database"
	"github.com/darthsalad/univboard/pkg/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
			logger.Fatalf("err loading: %v", err)
	}

	db, err := database.Connect(os.Getenv("DSN"))
	if err != nil {
		logger.Fatalf("err connecting: %v", err)
	}
	db.Init()
	
	defer db.Close()

	server := server.CreateServer(db)

	go func() {
		if err := server.Start("localhost:8000"); err != nil {
			logger.Fatalf("err starting server: %v", err)
		}
	}()

	logger.Logln("HTTP Server started on port 8000!")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Stop(ctx); err != nil {
		logger.Fatalf("err stopping server: %v", err)
	}
}