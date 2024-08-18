package app

import (
	"apartment_search_service/internal/config"
	"apartment_search_service/internal/handlers"
	"apartment_search_service/internal/logger"
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type App struct {
	serviceProvider *serviceProvider
	config          *config.Config
	logger          *logrus.Logger
	router          *mux.Router
}

func NewApp(ctx context.Context) *App {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		a.logger.Fatal("failed to init app: %s", err.Error())
		return nil
	}
	return a
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initLogger,
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

func (a *App) Run() {
	a.runServer()
}

func (a *App) initLogger(_ context.Context) error {
	a.logger = logger.NewLogger()
	return nil
}

func (a *App) initServer(_ context.Context) error {
	a.router = handlers.SetupRoutes(
		a.serviceProvider.userHandler,
		a.serviceProvider.houseHandler,
		a.serviceProvider.flatHandler)
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	a.config = config.LoadConfig()
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return a.serviceProvider.initServices(a.config, a.logger)
}

func (a *App) runServer() {
	a.logger.Printf("connect to http://localhost:%s/", a.config.Port)
	a.logger.Fatal(http.ListenAndServe(":"+a.config.Port, a.router))
}
