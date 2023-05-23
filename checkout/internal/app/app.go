package app

import (
	"context"
	"net"
	"net/http"
	"route256/checkout/internal/config"
	"route256/checkout/internal/handlers"
	"route256/checkout/internal/log"
	"time"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
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
	if err := app.RunHTTPServer(); err != nil {
		log.Fatalf("ERR: ", err)
	}

	return nil
}

func (app *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		config.Init,
		log.InitLogger,
		app.initServiceProvider,
		app.initHTTPServer,
	}

	for _, init := range inits {
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

func (app *App) initHTTPServer(_ context.Context) error {
	const timeout = 15

	app.httpServer = &http.Server{
		Addr:         net.JoinHostPort(config.AppConfig.Host, config.AppConfig.Port),
		Handler:      handlers.InitRouter(app.serviceProvider.cartService),
		ReadTimeout:  timeout * time.Second,
		WriteTimeout: timeout * time.Second,
	}

	return nil
}

func (app *App) RunHTTPServer() error {
	log.Infof("Start: HTTP server listening on port %s", config.AppConfig.Port)

	if err := app.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
