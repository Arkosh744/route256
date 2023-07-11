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
	"route256/libs/closer"
	"route256/libs/interceptor"
	"route256/libs/log"
	"route256/libs/metrics"
	"route256/libs/tracing"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	descCheckoutV1 "route256/pkg/checkout_v1"
	_ "route256/pkg/statik"

	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider  *serviceProvider
	grpcServer       *grpc.Server
	httpServer       *http.Server
	swaggerServer    *http.Server
	prometheusServer *http.Server
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
			log.Fatal("failed to run grpc server", zap.Error(err))
		}
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		err := app.RunHTTPServer()
		if err != nil {
			log.Fatal("failed to run http server", zap.Error(err))
		}
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		err := app.RunSwaggerServer()
		if err != nil {
			log.Fatal("failed to run swagger server", zap.Error(err))
		}
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		err := app.RunPrometheusServer()
		if err != nil {
			log.Fatal("failed to run prometheus server", zap.Error(err))
		}
	}()

	wg.Wait()

	return nil
}

func (app *App) initDeps(ctx context.Context) error {
	for _, init := range []func(context.Context) error{
		config.Init,
		metrics.Init,
		app.initLogger,
		app.initJaeger,
		app.initServiceProvider,
		app.initGrpcServer,
		app.initHTTPServer,
		app.initSwaggerServer,
		app.initPrometheusServer,
	} {
		if err := init(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (app *App) initLogger(ctx context.Context) error {
	if err := log.InitLogger(ctx, config.AppConfig.Log.Preset); err != nil {
		return err
	}

	return nil
}

func (app *App) initJaeger(_ context.Context) error {
	if err := tracing.Init(config.AppConfig.GetJaegerAddr(), "checkout"); err != nil {
		return err
	}

	return nil
}

func (app *App) initServiceProvider(ctx context.Context) error {
	app.serviceProvider = newServiceProvider(ctx)

	return nil
}

func (app *App) initGrpcServer(ctx context.Context) error {
	app.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer()),
			interceptor.LoggingInterceptor,
			interceptor.ValidateInterceptor,
		),
		),
	)
	reflection.Register(app.grpcServer)

	descCheckoutV1.RegisterCheckoutServer(app.grpcServer, app.serviceProvider.GetCheckoutImpl(ctx))

	return nil
}

func (app *App) RunGrpcServer() error {
	log.Info(fmt.Sprintf("GRPC server listening on port %s", config.AppConfig.GetGRPCAddr()))

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
	log.Info(fmt.Sprintf("Start: HTTP server listening on port %s", config.AppConfig.GetHTTPAddr()))

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
	log.Info(fmt.Sprintf("Swagger server is running on %s", config.AppConfig.GetSwaggerAddr()))

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

func (app *App) initPrometheusServer(_ context.Context) error {
	metricsServer := &http.Server{
		Addr: config.AppConfig.GetMetricsAddr(),
	}

	http.Handle("/metrics", promhttp.Handler())

	app.prometheusServer = metricsServer

	return nil
}

func (app *App) RunPrometheusServer() error {
	log.Info(fmt.Sprintf("Prometheus server is running on %s", config.AppConfig.GetMetricsAddr()))

	err := app.prometheusServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
