package task

import (
	"github.com/hoaxoan/onef-api/onef_task/delivery/http"
	"github.com/hoaxoan/onef-api/onef_task/repository"
	"github.com/hoaxoan/onef-api/onef_task/usecase"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewHandler(e *echo.Echo, client *mongo.Client) {
	repo := repository.NewRepository(client)
	usecase := usecase.NewUsecase(repo)
	http.NewHandler(e, usecase)
}
