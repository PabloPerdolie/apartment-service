package app

import (
	"apartment_search_service/internal/config"
	"apartment_search_service/internal/handlers"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	serviceProvider *serviceProvider
	config          *config.Config
	router          *mux.Router
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) Run() error {
	return a.runServer()
}

func (a *App) initServer(ctx context.Context) error {
	a.router = handlers.SetupRoutes(a.serviceProvider.userHandler)
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	a.config = config.LoadConfig()
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return a.serviceProvider.initServices(a.config)
}

func (a *App) runServer() error {
	log.Printf("connect to http://localhost:%s/", a.config.Port)
	log.Fatal(http.ListenAndServe(":"+a.config.Port, a.router))

	return nil
}
