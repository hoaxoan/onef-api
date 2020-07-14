package posts

import (
	"github.com/hoaxoan/onef-api/onef_posts/delivery/http"
	"github.com/hoaxoan/onef-api/onef_posts/repository"
	"github.com/hoaxoan/onef-api/onef_posts/usecase"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func NewHandler(e *echo.Echo, db *gorm.DB) {
	// Post
	postRepo := repository.NewPostRepository(db)
	postUsecase := usecase.NewPostUsecase(postRepo)
	http.NewPostHandler(e, postUsecase)

	// Posts
	postsRepo := repository.NewPostsRepository(db, postRepo)
	postsUsecase := usecase.NewPostsUsecase(postsRepo)
	http.NewPostsHandler(e, postsUsecase)
}
