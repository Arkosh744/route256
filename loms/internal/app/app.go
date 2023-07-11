package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"

	"route256/libs/closer"
	"route256/libs/metrics"
	"route256/libs/tracing"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"route256/libs/interceptor"
	"route256/libs/log"
	"route256/loms/internal/config"
	descLomsV1 "route256/pkg/loms_v1"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider  *serviceProvider
	grpcServer       *grpc.Server
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

	wg := &sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()

		if err := app.RunGrpcServer(); err != nil {
			log.Fatal("failed to run grpc server", zap.Error(err))
		}
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		if err := app.RunPrometheusServer(); err != nil {
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
		app.initPrometheusServer,
	} {
		if err := init(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (app *App) initJaeger(_ context.Context) error {
	if err := tracing.Init(config.AppConfig.GetJaegerAddr(), "loms"); err != nil {
		return err
	}

	return nil
}

func (app *App) initLogger(ctx context.Context) error {
	if err := log.InitLogger(ctx, config.AppConfig.Log.Preset); err != nil {
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
		)),
	)
	reflection.Register(app.grpcServer)

	descLomsV1.RegisterLomsServer(app.grpcServer, app.serviceProvider.GetLomsImpl(ctx))

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
