package categories

import (
	"github.com/hoaxoan/onef-api/onef_categories/delivery/http"
	"github.com/hoaxoan/onef-api/onef_categories/repository"
	"github.com/hoaxoan/onef-api/onef_categories/usecase"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func NewHandler(e *echo.Echo, db *gorm.DB) {
	repo := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repo)
	http.NewHandler(e, usecase)
}
