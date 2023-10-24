package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Din4EE/note-service-api/internal/app/api/note_v1"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	httpServer      *http.Server
	httpMux         *runtime.ServeMux
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
	defer func() {
		log.Println("Closing DB")
		if err := a.serviceProvider.GetDB(ctx).Close(); err != nil {
			log.Printf("failed to close db: %s", err.Error())
		}
	}()
	defer func() {
		log.Println("Stopping GRPC server")
		a.stopGRPCServer()
	}()
	defer func() {
		log.Println("Stopping HTTP server")
		if err := a.stopHTTPServer(ctx); err != nil {
			log.Printf("failed to stop http server: %s", err.Error())
		}
	}()

	grpcErrorCh := make(chan error)
	httpErrorCh := make(chan error)

	go func() {
		if err := a.runGRPCServer(); err != nil {
			grpcErrorCh <- err
		}
	}()

	go func() {
		if err := a.runHTTPServer(); err != nil {
			httpErrorCh <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-quit:
		log.Println("Graceful shutdown...")
	case err := <-grpcErrorCh:
		return fmt.Errorf("failed to run grpc server: %w", err)
	case err := <-httpErrorCh:
		return fmt.Errorf("failed to run http server: %w", err)
	}

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initService,
		a.initGRPCServer,
		a.initHTTPHandlers,
		a.initHTTPServer,
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

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()))
	desc.RegisterNoteServiceServer(a.grpcServer, a.noteImpl)

	return nil
}

func (a *App) initService(ctx context.Context) error {
	a.noteImpl = note_v1.NewNote(a.serviceProvider.GetNoteService(ctx))

	return nil
}

func (a *App) initHTTPHandlers(ctx context.Context) error {
	a.httpMux = runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := desc.RegisterNoteServiceHandlerFromEndpoint(ctx, a.httpMux, net.JoinHostPort(a.serviceProvider.GetConfig().GRPC.Host, a.serviceProvider.GetConfig().GRPC.Port), opts)
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

func (a *App) stopGRPCServer() {
	if a.grpcServer != nil {
		a.grpcServer.GracefulStop()
	}
}

func (a *App) initHTTPServer(_ context.Context) error {
	a.httpServer = &http.Server{
		Addr:    net.JoinHostPort(a.serviceProvider.GetConfig().HTTP.Host, a.serviceProvider.GetConfig().HTTP.Port),
		Handler: a.httpMux,
	}

	return nil
}

func (a *App) runHTTPServer() error {
	if err := a.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (a *App) stopHTTPServer(ctx context.Context) error {
	if a.httpServer == nil {
		return fmt.Errorf("HTTP server is not initialized")
	}

	return a.httpServer.Shutdown(ctx)
}
