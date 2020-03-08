package main

import (
	"log"

	"github.com/hoaxoan/onef-api/onef_core/repository"
	"github.com/hoaxoan/onef-api/onef_handler/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Recover())

	client, err := repository.Connection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	auth.NewHandler(e, client)

	log.Println(e.Start(":9090"))
}
