package auth

import (
	"github.com/hoaxoan/onef-api/onef_auth/delivery/http"
	"github.com/hoaxoan/onef-api/onef_auth/repository"
	"github.com/hoaxoan/onef-api/onef_auth/service"
	"github.com/hoaxoan/onef-api/onef_auth/usecase"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func NewHandler(e *echo.Echo, db *gorm.DB) {
	repo := repository.NewRepository(db)
	tokenService := service.NewTokeService(repo)
	usecase := usecase.NewUsecase(repo, tokenService)
	http.NewHandler(e, usecase)
}
