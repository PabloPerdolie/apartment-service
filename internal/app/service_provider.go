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

func (sp *serviceProvider) initServices(conf *config.Config) error {
	db, err := utils.InitDB(conf.DB.User, conf.DB.Password, conf.DB.Name, conf.DB.Host, conf.DB.Port)
	if err != nil {
		return err
	}
	sp.userRepo = userrepo.NewUserRepository(db)
	sp.authService = auth.NewAuthService(sp.userRepo)
	sp.userHandler = user.NewHandler(sp.authService)

	sp.houseRepo = house2.NewHouseRepository(db)
	sp.houseService = house3.NewHouseService(sp.houseRepo, sp.flatRepo)
	sp.houseHandler = house.NewHandler(sp.houseService)

	sp.flatRepo = flat2.NewFlatRepository(db)
	sp.flatService = flat3.NewFlatService(sp.flatRepo, sp.houseRepo)
	sp.flatHandler = flat.NewHandler(sp.flatService)

	return nil
}
