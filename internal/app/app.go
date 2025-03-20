package app

import (
	"app/internal/api"
	"app/internal/config"
	"app/internal/database"
	"app/internal/entity"
	"app/internal/service/users"
	"net/http"
)

// Application  ...
type Application struct {
	BindAddr string
}

func (a *Application) Run() error {
	configLoader := config.Loader{}

	cfg, err := configLoader.LoadConfig()
	if err != nil {
		return err
	}

	db, err := database.ConnectAndMigrate(cfg)

	if err != nil {
		return err
	}

	repo := entity.NewRepository(db)

	securityUseCase := users.SecurityUseCase{SecurityRepository: repo}
	getUsersListUseCase := users.GetUsersListUseCase{GetUsersRepository: repo}
	getUserUseCase := users.GetUserUseCase{GetUserRepository: repo}
	createUserUseCase := users.CreateUserUseCase{CreateUserRepository: repo}
	updateUserUseCase := users.UpdateUserUseCase{UpdateUserRepository: repo}
	removeUserUseCase := users.RemoveUserUseCase{RemoveUserRepository: repo}

	s := api.New(api.ServerParams{
		Config:              cfg,
		SecurityUseCase:     securityUseCase,
		GetUserUseCase:      getUserUseCase,
		GetUsersListUseCase: getUsersListUseCase,
		CreateUserUseCase:   createUserUseCase,
		UpdateUserUseCase:   updateUserUseCase,
		RemoveUserUseCase:   removeUserUseCase,
	})

	return http.ListenAndServe(a.BindAddr, s.Router)
}

func New(bindAddr string) *Application {
	return &Application{BindAddr: bindAddr}
}
