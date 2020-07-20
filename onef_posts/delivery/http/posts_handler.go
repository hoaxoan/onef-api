package http

import (
	"net/http"
	"strconv"

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

	g.GET("/{postUuid}/comments/{postCommentId}", handler.GetPostComment)
	g.GET("/{postUuid}/comments", handler.GetCommentsForPostWithUuid)
	g.GET("/{postUuid}/comments/{postCommentId}/replies", handler.GetRepliesForCommentWithIdForPostWithUuid)
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

	count, err := strconv.ParseInt(ctx.QueryParam("count"), 10, 64)
	if err == nil {
		req.Count = count
	}

	maxId, err := strconv.ParseInt(ctx.QueryParam("max_id"), 10, 64)
	if err == nil {
		req.MaxId = maxId
	} else {
		req.MaxId = 0
	}

	req.UserName = ctx.QueryParam("username")

	var res model.PostResponse
	if err := h.UUcase.GetTimelinePosts(ctx.Request().Context(), &req, &res); err != nil {
		return ctx.JSON(http.StatusNotFound, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.Posts)
}

func (h *postsHandler) GetPostComment(ctx echo.Context) error {
	postUuid := ctx.QueryParam("postUuid")
	postCommentId, err := strconv.ParseInt(ctx.QueryParam("postCommentId"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.PostCommentResponse
	if err := h.UUcase.GetPostComment(ctx.Request().Context(), postUuid, postCommentId, &res); err != nil {
		return ctx.JSON(http.StatusNotFound, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.PostComment)
}

func (h *postsHandler) GetCommentsForPostWithUuid(ctx echo.Context) error {
	var req model.PostCommentRequest

	count, err := strconv.ParseInt(ctx.QueryParam("count"), 10, 64)
	if err == nil {
		req.Count = count
	}

	maxId, err := strconv.ParseInt(ctx.QueryParam("max_id"), 10, 64)
	if err == nil {
		req.MaxId = maxId
	} else {
		req.MaxId = 0
	}

	req.UserName = ctx.QueryParam("username")

	var res model.PostCommentResponse
	if err := h.UUcase.GetCommentsForPostWithUuid(ctx.Request().Context(), &req, &res); err != nil {
		return ctx.JSON(http.StatusNotFound, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.PostComment)
}

func (h *postsHandler) GetRepliesForCommentWithIdForPostWithUuid(ctx echo.Context) error {
	var req model.PostCommentRequest

	count, err := strconv.ParseInt(ctx.QueryParam("count"), 10, 64)
	if err == nil {
		req.Count = count
	}

	maxId, err := strconv.ParseInt(ctx.QueryParam("max_id"), 10, 64)
	if err == nil {
		req.MaxId = maxId
	} else {
		req.MaxId = 0
	}

	req.UserName = ctx.QueryParam("username")

	var res model.PostCommentResponse
	if err := h.UUcase.GetRepliesForCommentWithIdForPostWithUuid(ctx.Request().Context(), &req, &res); err != nil {
		return ctx.JSON(http.StatusNotFound, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.PostComment)
}
