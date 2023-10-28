package app

import (
	"context"
	"log"

	"github.com/Din4EE/note-service-api/internal/config"
	"github.com/Din4EE/note-service-api/internal/pkg/db"
	"github.com/Din4EE/note-service-api/internal/repo"
	repoNote "github.com/Din4EE/note-service-api/internal/repo/note"
	"github.com/Din4EE/note-service-api/internal/service/note"
)

type serviceProvider struct {
	noteService *note.Service
	noteRepo    repo.NoteRepository
	db          db.Client
	config      *config.Config
	configPath  string
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{
		configPath: configPath,
	}
}

func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		cfg, err := s.GetConfig().GetDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}
		dbc, err := db.NewClient(ctx, cfg)
		if err != nil {
			log.Fatalf("failed to init db: %s", err.Error())
		}
		s.db = dbc
	}

	return s.db
}

func (s *serviceProvider) GetConfig() *config.Config {
	if s.config == nil {
		cfg, err := config.NewConfig(s.configPath)
		if err != nil {
			log.Fatalf("failed to init config: %s", err.Error())
		}
		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) GetNoteRepository(ctx context.Context) repo.NoteRepository {
	if s.noteRepo == nil {
		s.noteRepo = repoNote.NewRepository(s.GetDB(ctx))
	}

	return s.noteRepo
}

func (s *serviceProvider) GetNoteService(ctx context.Context) *note.Service {
	if s.noteService == nil {
		s.noteService = note.NewService(s.GetNoteRepository(ctx))
	}

	return s.noteService
}
