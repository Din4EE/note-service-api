package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/Din4EE/note-service-api/internal/app/api/note_v1"
	"github.com/Din4EE/note-service-api/internal/config"
	"github.com/Din4EE/note-service-api/internal/repo/note"
	"github.com/Din4EE/note-service-api/internal/service/note"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load env: %s", err.Error())
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		log.Fatal(startGRPC())
	}()

	go func() {
		defer wg.Done()
		log.Fatal(startHTTP(ctx))
	}()

	wg.Wait()
}

func startGRPC() error {
	list, err := net.Listen("tcp", os.Getenv("GRPC_SERVER_HOST")+":"+os.Getenv("GRPC_SERVER_PORT"))
	if err != nil {
		return err
	}

	noteService := note.NewService(repo.NewPostgresNoteRepository(config.PgConfig{
		DSN: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_SSLMODE")),
	}))
	s := grpc.NewServer(grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()))
	desc.RegisterNoteServiceServer(s, note_v1.NewNote(noteService))

	if err = s.Serve(list); err != nil {
		return err
	}

	return nil
}

func startHTTP(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := desc.RegisterNoteServiceHandlerFromEndpoint(ctx, mux, os.Getenv("GRPC_SERVER_HOST")+":"+os.Getenv("GRPC_SERVER_PORT"), opts)
	if err != nil {
		return err
	}

	if err = http.ListenAndServe(os.Getenv("HTTP_SERVER_HOST")+":"+os.Getenv("HTTP_SERVER_PORT"), mux); err != nil {
		return err
	}

	return nil
}
