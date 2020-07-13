package main

import (
	"log"

	"github.com/hoaxoan/onef-api/onef_core/repository"
	"github.com/hoaxoan/onef-api/onef_handler/auth"
	"github.com/hoaxoan/onef-api/onef_handler/categories"
	"github.com/hoaxoan/onef-api/onef_handler/circles"
	"github.com/hoaxoan/onef-api/onef_handler/communities"
	"github.com/hoaxoan/onef-api/onef_handler/connections"
	"github.com/hoaxoan/onef-api/onef_handler/devices"
	"github.com/hoaxoan/onef-api/onef_handler/hashtags"
	"github.com/hoaxoan/onef-api/onef_handler/posts"
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
	categories.NewHandler(e, client)
	circles.NewHandler(e, client)
	communities.NewHandler(e, client)
	connections.NewHandler(e, client)
	devices.NewHandler(e, client)
	hashtags.NewHandler(e, client)
	posts.NewHandler(e, client)
	//story.NewHandler(e, client)

	log.Println(e.Start(":9090"))
}
