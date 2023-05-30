package app

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"

	"route256/checkout/internal/config"
	"route256/checkout/internal/log"
	"route256/libs/closer"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	descCheckoutV1 "route256/pkg/checkout_v1"
	_ "route256/pkg/statik"

	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	httpServer      *http.Server
	swaggerServer   *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()

		err := app.RunGrpcServer()
		if err != nil {
			log.Fatalf("failed to run grpc server: %v", err)
		}
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		err := app.RunHTTPServer()
		if err != nil {
			log.Fatalf("failed to run http server: %v", err)
		}
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		err := app.RunSwaggerServer()
		if err != nil {
			log.Fatalf("failed to run swagger server: %v", err)
		}
	}()

	wg.Wait()

	return nil
}

func (app *App) initDeps(ctx context.Context) error {
	for _, init := range []func(context.Context) error{
		config.Init,
		log.InitLogger,
		app.initServiceProvider,
		app.initGrpcServer,
		app.initHTTPServer,
		app.initSwaggerServer,
	} {
		if err := init(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (app *App) initServiceProvider(ctx context.Context) error {
	app.serviceProvider = newServiceProvider(ctx)

	return nil
}

func (app *App) initGrpcServer(ctx context.Context) error {
	app.grpcServer = grpc.NewServer()
	reflection.Register(app.grpcServer)

	descCheckoutV1.RegisterCheckoutServer(app.grpcServer, app.serviceProvider.GetCheckoutImpl(ctx))

	return nil
}

func (app *App) RunGrpcServer() error {
	log.Infof("GRPC server listening on port %s", config.AppConfig.GetGRPCAddr())

	list, err := net.Listen("tcp", config.AppConfig.GetGRPCAddr())
	if err != nil {
		return err
	}

	err = app.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) initHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := descCheckoutV1.RegisterCheckoutHandlerFromEndpoint(ctx, mux, config.AppConfig.GetGRPCAddr(), opts)
	if err != nil {
		return err
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	timeout, err := time.ParseDuration(config.AppConfig.Timeout)
	if err != nil {
		return fmt.Errorf("failed to parse timeout: %w", err)
	}

	app.httpServer = &http.Server{
		Addr:              config.AppConfig.GetHTTPAddr(),
		Handler:           corsMiddleware.Handler(mux),
		ReadHeaderTimeout: timeout,
	}

	return nil
}

func (app *App) RunHTTPServer() error {
	log.Infof("Start: HTTP server listening on port %s", config.AppConfig.GetHTTPAddr())

	if err := app.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (app *App) initSwaggerServer(_ context.Context) error {
	statikFs, err := fs.New()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(statikFs)))
	mux.HandleFunc("/swagger.json", serveSwaggerFile("/swagger.json"))

	timeout, err := time.ParseDuration(config.AppConfig.Timeout)
	if err != nil {
		return fmt.Errorf("failed to parse timeout: %w", err)
	}

	app.swaggerServer = &http.Server{
		Addr:              config.AppConfig.GetSwaggerAddr(),
		Handler:           mux,
		ReadHeaderTimeout: timeout,
	}

	return nil
}

func (app *App) RunSwaggerServer() error {
	log.Infof("Swagger server is running on %s", config.AppConfig.GetSwaggerAddr())

	err := app.swaggerServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func serveSwaggerFile(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statikFs, err := fs.New()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		file, err := statikFs.Open(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		w.Header().Set("Content-Type", "application/json")

		_, err = w.Write(content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	}
}
