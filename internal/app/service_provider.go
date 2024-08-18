package app

import (
	"apartment_search_service/internal/config"
	"apartment_search_service/internal/handlers/flat"
	"apartment_search_service/internal/handlers/house"
	"apartment_search_service/internal/handlers/user"
	"apartment_search_service/internal/repositories"
	flat2 "apartment_search_service/internal/repositories/flat"
	house2 "apartment_search_service/internal/repositories/house"
	userrepo "apartment_search_service/internal/repositories/user"
	"apartment_search_service/internal/services"
	"apartment_search_service/internal/services/auth"
	flat3 "apartment_search_service/internal/services/flat"
	house3 "apartment_search_service/internal/services/house"
	"apartment_search_service/internal/utils"
	"github.com/sirupsen/logrus"
)

type serviceProvider struct {
	userHandler  *user.Handler
	authService  services.AuthService
	userRepo     repositories.UserRepository
	houseHandler *house.Handler
	houseService services.HouseService
	houseRepo    repositories.HouseRepository
	flatHandler  *flat.Handler
	flatService  services.FlatService
	flatRepo     repositories.FlatRepository
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) initServices(conf *config.Config, logger *logrus.Logger) error {
	db, err := utils.InitDB(conf.DB.User, conf.DB.Password, conf.DB.Name, conf.DB.Host, conf.DB.Port, logger)
	if err != nil {
		return err
	}
	sp.userRepo = userrepo.NewUserRepository(db)
	sp.flatRepo = flat2.NewFlatRepository(db)
	sp.houseRepo = house2.NewHouseRepository(db)

	sp.authService = auth.NewAuthService(sp.userRepo)
	sp.houseService = house3.NewHouseService(sp.houseRepo, sp.flatRepo)
	sp.flatService = flat3.NewFlatService(sp.flatRepo, sp.houseRepo)

	sp.userHandler = user.NewHandler(sp.authService, logger)
	sp.houseHandler = house.NewHandler(sp.houseService, logger)
	sp.flatHandler = flat.NewHandler(sp.flatService, logger)

	return nil
}
