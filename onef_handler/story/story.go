package story

import (
	"github.com/hoaxoan/onef-api/onef_story/delivery/http"
	"github.com/hoaxoan/onef-api/onef_story/repository"
	"github.com/hoaxoan/onef-api/onef_story/usecase"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewHandler(e *echo.Echo, client *mongo.Client) {
	repo := repository.NewRepository(client)
	usecase := usecase.NewUsecase(repo)
	http.NewHandler(e, usecase)
}
