package http

import (
	"net/http"

	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_posts"

	"github.com/labstack/echo/v4"
)

type postsHandler struct {
	UUcase onef_posts.PostsUsecase
}

func NewPostsHandler(e *echo.Echo, uc onef_posts.PostsUsecase) {
	handler := &postsHandler{
		UUcase: uc,
	}

	PublicPostsRoute(e, handler)
}

func PublicPostsRoute(e *echo.Echo, handler *postsHandler) {
	g := e.Group("/api/v1/posts")
	g.GET("/top", handler.GetTopPosts)
	g.GET("/trending", handler.GetTrendingPosts)
	g.GET("", handler.GetTimelinePosts)
}

func (h *postsHandler) GetTopPosts(ctx echo.Context) error {
	var req model.TopPostRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.TopPostResponse
	if err := h.UUcase.GetTopPosts(ctx.Request().Context(), &req, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.TopPosts)
}

func (h *postsHandler) GetTrendingPosts(ctx echo.Context) error {
	var req model.TrendingPostRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.TrendingPostResponse
	if err := h.UUcase.GetTrendingPosts(ctx.Request().Context(), &req, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.TrendingPosts)
}

func (h *postsHandler) GetTimelinePosts(ctx echo.Context) error {
	var req model.PostRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.PostResponse
	if err := h.UUcase.GetTimelinePosts(ctx.Request().Context(), &req, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.Posts)
}
