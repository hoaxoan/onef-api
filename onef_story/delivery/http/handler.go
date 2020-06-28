package http

import (
	"net/http"

	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_story"

	"github.com/labstack/echo/v4"
)

type storyHandler struct {
	UUcase onef_story.Usecase
}

func NewHandler(e *echo.Echo, uc onef_story.Usecase) {
	handler := &storyHandler{
		UUcase: uc,
	}

	PublicRoute(e, handler)
}

func PublicRoute(e *echo.Echo, handler *storyHandler) {
	g := e.Group("/api/v1/story")
	g.PATCH("/stories", handler.GetAll)
	g.PATCH("/search", handler.Search)
	g.PATCH("/story", handler.Get)
	g.POST("/story", handler.Create)
	g.PUT("/story", handler.Update)
	g.DELETE("/story", handler.Delete)
}

func (h *storyHandler) GetAll(ctx echo.Context) error {
	var req model.StoryRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.StoryResponse
	err := h.UUcase.GetAll(ctx.Request().Context(), &req, &res)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Stories)
}

func (h *storyHandler) Search(ctx echo.Context) error {
	var req model.StoryRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.StoryResponse
	err := h.UUcase.GetAll(ctx.Request().Context(), &req, &res)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Stories)
}

func (h *storyHandler) Get(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	var res model.StoryResponse
	if err := h.UUcase.Get(ctx.Request().Context(), id, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Stories)
}

func (h *storyHandler) Create(ctx echo.Context) error {
	var story model.Story
	if err := ctx.Bind(&story); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.StoryResponse
	if err := h.UUcase.Create(ctx.Request().Context(), &story, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.Story)
}

func (h *storyHandler) Update(ctx echo.Context) error {
	var story model.Story
	if err := ctx.Bind(&story); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.StoryResponse
	if err := h.UUcase.Update(ctx.Request().Context(), &story, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.Story)
}

func (h *storyHandler) Delete(ctx echo.Context) error {
	var story model.Story
	if err := ctx.Bind(&story); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.StoryResponse
	if err := h.UUcase.Delete(ctx.Request().Context(), &story, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.Story)
}
