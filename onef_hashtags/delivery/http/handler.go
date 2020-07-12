package http

import (
	"net/http"
	"strconv"

	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_hashtags"

	"github.com/labstack/echo/v4"
)

type handler struct {
	UUcase onef_hashtags.Usecase
}

func NewHandler(e *echo.Echo, uc onef_hashtags.Usecase) {
	handler := &handler{
		UUcase: uc,
	}

	PublicRoute(e, handler)
}

func PublicRoute(e *echo.Echo, handler *handler) {
	g := e.Group("/api/v1/hashtags")
	g.GET("/search", handler.Search)
	g.GET("/:id", handler.GetWithId)
	g.GET("", handler.Get)
	g.POST("", handler.Create)
	g.PUT("", handler.Update)
	g.DELETE("", handler.Delete)
}

func (h *handler) Search(ctx echo.Context) error {
	var req model.HashtagRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.HashtagResponse
	if err := h.UUcase.Get(ctx.Request().Context(), &req, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.Hashtags)
}

func (h *handler) Get(ctx echo.Context) error {
	var req model.HashtagRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.HashtagResponse
	if err := h.UUcase.Get(ctx.Request().Context(), &req, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.Hashtags)
}

func (h *handler) GetWithId(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.HashtagResponse
	if err := h.UUcase.GetWithId(ctx.Request().Context(), int64(id), &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.Hashtag)
}

func (h *handler) Create(ctx echo.Context) error {
	var hashtag model.Hashtag
	if err := ctx.Bind(&hashtag); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.HashtagResponse
	if err := h.UUcase.Create(ctx.Request().Context(), &hashtag, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Hashtag)
}

func (h *handler) Update(ctx echo.Context) error {
	var hashtag model.Hashtag
	if err := ctx.Bind(&hashtag); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.HashtagResponse
	if err := h.UUcase.Update(ctx.Request().Context(), &hashtag, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Hashtag)
}

func (h *handler) Delete(ctx echo.Context) error {
	var hashtag model.Hashtag
	if err := ctx.Bind(&hashtag); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.HashtagResponse
	if err := h.UUcase.Delete(ctx.Request().Context(), &hashtag, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Hashtag)
}
