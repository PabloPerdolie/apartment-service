package app

import (
	"apartment_search_service/internal/config"
	"apartment_search_service/internal/handlers/user"
	"apartment_search_service/internal/repositories"
	userrepo "apartment_search_service/internal/repositories/user"
	"apartment_search_service/internal/services"
	"apartment_search_service/internal/utils"
)

type serviceProvider struct {
	userHandler *user.Handler
	authService services.AuthService
	userRepo    repositories.UserRepository
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
	sp.authService = services.NewAuthService(sp.userRepo)
	sp.userHandler = user.NewHandler(sp.authService)

	return nil
}
