package main

import (
	"database/sql"
	"github.com/andreluizmicro/desafio-backend/config"
	"github.com/andreluizmicro/desafio-backend/internal/Application/user"
	"github.com/andreluizmicro/desafio-backend/internal/infrastructure/http"
	"github.com/andreluizmicro/desafio-backend/internal/infrastructure/http/controller"
	"github.com/andreluizmicro/desafio-backend/internal/infrastructure/repository"
	"github.com/andreluizmicro/desafio-backend/pkg/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg, err := config.LoadConfig("../")
	if err != nil {
		panic(err)
	}

	db, err := database.NewConnection(cfg)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	userRepository := repository.NewUserRepository(db)
	createUserService := user.NewCreateUserService(userRepository)
	userController := controller.NewUserController(createUserService)
	http.InitRoutes(userController, cfg.WebServerPort)
}
