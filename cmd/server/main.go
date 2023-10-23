package main

import (
	"context"
	"log"
	"os"

	"github.com/Din4EE/note-service-api/internal/app"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load env: %s", err.Error())
	}
}

func main() {
	ctx := context.Background()
	a, err := app.NewApp(ctx, os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	if err = a.Run(ctx); err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
