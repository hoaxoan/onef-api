package auth

import (
	"github.com/hoaxoan/onef-api/onef_auth/delivery/http"
	"github.com/hoaxoan/onef-api/onef_auth/repository"
	"github.com/hoaxoan/onef-api/onef_auth/service"
	"github.com/hoaxoan/onef-api/onef_auth/usecase"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewHandler(e *echo.Echo, client *mongo.Client) {
	repo := repository.NewUserRepository(client)
	tokenService := service.NewTokeService(repo)
	usecase := usecase.NewUsecase(repo, tokenService)
	http.NewHandler(e, usecase)
}
