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
	g.PATCH("/story", handler.Create)
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
	return ctx.JSON(http.StatusOK, res)
}
