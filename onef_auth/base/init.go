package base

import (
	"github.com/hoaxoan/onef-api/onef_auth/http"
	"github.com/hoaxoan/onef-api/onef_auth/repository"
	"github.com/hoaxoan/onef-api/onef_auth/service"
	"github.com/hoaxoan/onef-api/onef_auth/usecase"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewHandler(e *echo.Echo, Client *mongo.Client) {
	repo := repository.NewUserRepository(client)
	tokenService := service.NewTokeService(repo)
	usecase := usecase.NewUserUsecase(repo, tokenService)
	http.NewUserHandler(e, usecase)

}
