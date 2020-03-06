package main

import (
	"fmt"
	"log"

	"github.com/hoaxoan/onef-api/onef_auth/base"
	"github.com/hoaxoan/onef-api/onef_core/repository"
	"github.com/hoaxoan/onef-api/onef_core/setting"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Printf("config app: %+v", setting.Config)

	e := echo.New()
	e.Use(middleware.Recover())
	client, err := repository.Connection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	base.NewHandler(client)

	log.Println(e.Start(":9090"))
}
