package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_core/setting"
	"github.com/hoaxoan/onef-api/onef_posts"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type postHandler struct {
	UUcase onef_posts.PostUsecase
}

func NewPostHandler(e *echo.Echo, uc onef_posts.PostUsecase) {
	handler := &postHandler{
		UUcase: uc,
	}

	PublicPostRoute(e, handler)
}

func PublicPostRoute(e *echo.Echo, handler *postHandler) {
	JWTConfig := middleware.JWTConfig{
		SigningKey: []byte(setting.Config.JWTSecret.JWTKey),
		Claims:     &model.CustomClaims{},
	}

	g := e.Group("/api/v1/posts")
	g.Use(middleware.JWTWithConfig(JWTConfig))
	g.GET("/:postUuid", handler.GetPostWithUuid)
	g.GET("/:postUuid/status", handler.GetPostWithUuidStatus)
	g.POST("", handler.CreatePost)
	g.PUT("/:postUuid", handler.EditPost)
	g.POST("/:postUuid/publish", handler.PublishPost)
	g.DELETE("/:postUuid", handler.DeletePostWithUuid)

	g.PUT("", handler.Update)
	g.DELETE("", handler.Delete)
}

func (h *postHandler) GetWithId(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.PostResponse
	if err := h.UUcase.GetWithId(ctx.Request().Context(), int64(id), &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.Post)
}

func (h *postHandler) CreatePost(ctx echo.Context) error {
	userToken := ctx.Get("user").(*jwt.Token)
	if userToken == nil || userToken.Valid == false {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "token invalid"})
	}
	claims := userToken.Claims.(*model.CustomClaims)
	if claims == nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "token invalid"})
	}
	user := claims.User
	if user == nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "token invalid"})
	}

	var createPost model.CreatePostRequest
	if err := ctx.Bind(&createPost); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	createPost.Creator = user
	createPost.UUId = uuid.New().String()

	if createPost.CircleId != "" && len(strings.TrimSpace(createPost.CircleId)) > 0 {
		createPost.CircleIds = strings.Split(createPost.CircleId, ",")
	}

	var res model.PostResponse
	if err := h.UUcase.CreatePost(ctx.Request().Context(), &createPost, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, res.Post)
}

func (h *postHandler) GetPostWithUuid(ctx echo.Context) error {
	postUuid := ctx.Param("postUuid")

	var res model.PostResponse
	if err := h.UUcase.GetPostWithUuid(ctx.Request().Context(), postUuid, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.Post)
}

func (h *postHandler) GetPostWithUuidStatus(ctx echo.Context) error {
	postUuid := ctx.Param("postUuid")

	var res model.PostResponse
	if err := h.UUcase.GetPostWithUuid(ctx.Request().Context(), postUuid, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.Post)
}

func (h *postHandler) EditPost(ctx echo.Context) error {
	postUuid := ctx.Param("postUuid")

	var req model.EditPostRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.PostResponse
	if err := h.UUcase.GetPostWithUuid(ctx.Request().Context(), postUuid, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	post := res.Post
	post.Text = req.Text

	var response model.PostResponse
	if err := h.UUcase.Update(ctx.Request().Context(), post, &response); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, response.Post)
}

func (h *postHandler) PublishPost(ctx echo.Context) error {
	postUuid := ctx.Param("postUuid")
	var res model.PostResponse
	if err := h.UUcase.PublishPost(ctx.Request().Context(), postUuid, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, true)
}

func (h *postHandler) DeletePostWithUuid(ctx echo.Context) error {
	postUuid := ctx.Param("postUuid")

	var res model.PostResponse
	if err := h.UUcase.GetPostWithUuid(ctx.Request().Context(), postUuid, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var response model.PostResponse
	if err := h.UUcase.Delete(ctx.Request().Context(), res.Post, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, response.Post)
}

func (h *postHandler) Update(ctx echo.Context) error {
	var post model.Post
	if err := ctx.Bind(&post); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.PostResponse
	if err := h.UUcase.Update(ctx.Request().Context(), &post, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Post)
}

func (h *postHandler) Delete(ctx echo.Context) error {
	var post model.Post
	if err := ctx.Bind(&post); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.PostResponse
	if err := h.UUcase.Delete(ctx.Request().Context(), &post, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Post)
}
