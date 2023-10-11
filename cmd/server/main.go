package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Din4EE/note-service-api/config"
	"github.com/Din4EE/note-service-api/internal/app/api/note_v1"
	"github.com/Din4EE/note-service-api/internal/app/repo"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

const port = ":50051"

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load env: %s", err.Error())
	}
}

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to mapping port %s", err.Error())
	}

	r := repo.NewPostgresNoteRepository(config.PgConfig{
		DSN: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_SSLMODE")),
	})
	noteService := note_v1.NewNote(r)
	s := grpc.NewServer()
	desc.RegisterNoteServiceServer(s, noteService)
	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
