package main

import (
	"log"
	"net"

	"github.com/Din4EE/note-service-api/internal/app/api/note_v1"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"github.com/Din4EE/note-service-api/repo"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to mapping port %s", err.Error())
	}

	r := repo.NewInMemoryNoteRepository()
	noteService := note_v1.NewNote(r)
	s := grpc.NewServer()
	desc.RegisterNoteServiceServer(s, noteService)
	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
