package main

import (
	"fmt"
	"log"

	"github.com/hoaxoan/onef-api/onef_core/repository"
	"github.com/hoaxoan/onef-api/onef_core/setting"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Printf("config app: %+v", setting.Config)

	e := echo.New()
	e.Use(middleware.Recover())
	_, err := repository.Connection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// usRepo := _userRepository.NewUserRepository(client)
	// tokenService := _userSerice.NewTokeService(usRepo)
	// uu := _userUsecase.NewUserUsecase(usRepo, tokenService)
	// _userHttpDeliver.NewUserHandler(e, uu)

	log.Println(e.Start(":9090"))
}
