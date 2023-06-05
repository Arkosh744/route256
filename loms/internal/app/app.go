package app

import (
	"context"
	"net"

	"route256/libs/interceptor"
	"route256/libs/log"
	"route256/loms/internal/config"
	descLomsV1 "route256/pkg/loms_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
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
	err := app.RunGrpcServer()
	if err != nil {
		log.Fatalf("failed to run grpc server: %v", err)
	}

	return nil
}

func (app *App) initDeps(ctx context.Context) error {
	for _, init := range []func(context.Context) error{
		config.Init,
		app.initLogger,
		app.initServiceProvider,
		app.initGrpcServer,
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

func (app *App) initServiceProvider(ctx context.Context) error {
	app.serviceProvider = newServiceProvider(ctx)

	return nil
}

func (app *App) initGrpcServer(ctx context.Context) error {
	app.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor),
	)
	reflection.Register(app.grpcServer)

	descLomsV1.RegisterLomsServer(app.grpcServer, app.serviceProvider.GetLomsImpl(ctx))

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
