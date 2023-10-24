package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/Din4EE/note-service-api/internal/app/api/note_v1"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	httpServer      *runtime.ServeMux
	grpcServer      *grpc.Server
	noteImpl        *note_v1.Note
	serviceProvider *serviceProvider
	configPath      string
}

func NewApp(ctx context.Context, configPath string) (*App, error) {
	a := &App{
		configPath: configPath,
	}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		log.Fatal(a.runGRPCServer())
	}()

	go func() {
		defer wg.Done()
		log.Fatal(a.runHTTPServer())
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initService,
		a.initGRPCServer,
		a.initHTTPHandlers,
	}

	for _, init := range inits {
		if err := init(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.configPath)

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()))
	desc.RegisterNoteServiceServer(a.grpcServer, a.noteImpl)

	return nil
}

func (a *App) initService(ctx context.Context) error {
	a.noteImpl = note_v1.NewNote(a.serviceProvider.GetNoteService(ctx))

	return nil
}

func (a *App) initHTTPHandlers(ctx context.Context) error {
	a.httpServer = runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := desc.RegisterNoteServiceHandlerFromEndpoint(ctx, a.httpServer, net.JoinHostPort(a.serviceProvider.GetConfig().GRPC.Host, a.serviceProvider.GetConfig().GRPC.Port), opts)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runGRPCServer() error {
	list, err := net.Listen("tcp", net.JoinHostPort(a.serviceProvider.GetConfig().GRPC.Host, a.serviceProvider.GetConfig().GRPC.Port))
	if err != nil {
		return err
	}

	if err = a.grpcServer.Serve(list); err != nil {
		return err
	}

	return nil
}

func (a *App) runHTTPServer() error {
	if err := http.ListenAndServe(net.JoinHostPort(a.serviceProvider.GetConfig().HTTP.Host, a.serviceProvider.GetConfig().HTTP.Port), a.httpServer); err != nil {
		return err
	}

	return nil
}
