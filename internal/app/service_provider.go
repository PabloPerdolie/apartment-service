package app

import (
	"apartment_search_service/internal/config"
	"apartment_search_service/internal/handlers/flat"
	"apartment_search_service/internal/handlers/house"
	"apartment_search_service/internal/handlers/user"
	"apartment_search_service/internal/repositories"
	"apartment_search_service/internal/repositories/in_memory"
	"apartment_search_service/internal/repositories/postgres"
	"apartment_search_service/internal/services"
	"apartment_search_service/internal/utils"
	"apartment_search_service/internal/utils/sender"
	"github.com/sirupsen/logrus"
)

type serviceProvider struct {
	userHandler         *user.Handler
	authService         services.AuthService
	userRepo            repositories.UserRepository
	houseHandler        *house.Handler
	houseService        services.HouseService
	houseRepo           repositories.HouseRepository
	flatHandler         *flat.Handler
	flatService         services.FlatService
	flatRepo            repositories.FlatRepository
	subscriptionRepo    repositories.SubscriptionRepository
	subscriptionService services.SubscriptionService
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) initServices(conf *config.Config, logger *logrus.Logger) error {
	db, err := utils.InitDB(conf.DB.User, conf.DB.Password, conf.DB.Name, conf.DB.Host, conf.DB.Port, logger)
	if err != nil {
		return err
	}
	sp.userRepo = postgres.NewUserRepository(db)
	sp.flatRepo = postgres.NewFlatRepository(db)
	sp.houseRepo = postgres.NewHouseRepository(db)
	sp.subscriptionRepo = in_memory.NewSubscriptionRepo()

	sp.authService = services.NewAuthService(sp.userRepo)
	sp.houseService = services.NewHouseService(sp.houseRepo, sp.flatRepo)
	sp.flatService = services.NewFlatService(sp.flatRepo, sp.houseRepo)
	sp.subscriptionService = services.NewSubscriptionService(sp.houseRepo, sp.subscriptionRepo, sender.New(), logger)

	sp.userHandler = user.NewHandler(sp.authService, logger)
	sp.houseHandler = house.NewHandler(sp.houseService, sp.subscriptionService, logger)
	sp.flatHandler = flat.NewHandler(sp.flatService, sp.subscriptionService, logger)

	return nil
}
