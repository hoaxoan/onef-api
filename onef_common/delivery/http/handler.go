package http

import (
	"net/http"

	"github.com/hoaxoan/onef-api/onef_common"
	"github.com/hoaxoan/onef-api/onef_core/model"

	"github.com/labstack/echo/v4"
)

type commonHandler struct {
	UUcase onef_common.Usecase
}

func NewHandler(e *echo.Echo, uc onef_common.Usecase) {
	handler := &commonHandler{
		UUcase: uc,
	}

	PublicRoute(e, handler)
}

func PublicRoute(e *echo.Echo, handler *commonHandler) {
	g := e.Group("/api/v1/common")
	// Categories
	g.PATCH("/categories", handler.GetCategories)
	g.POST("/categories", handler.CreateCategories)
	g.POST("/category", handler.CreateCategory)
	g.PUT("/category", handler.UpdateCategory)
	g.DELETE("/category", handler.DeleteCategory)

	// Moods
	g.PATCH("/moods", handler.GetMoods)
	g.POST("/moods", handler.CreateMoods)
	g.POST("/mood", handler.CreateMood)
	g.PUT("/mood", handler.UpdateMood)
	g.DELETE("/mood", handler.DeleteMood)

	// Color Ranges
	g.PATCH("/color-ranges", handler.GetColorRanges)
	g.POST("/color-ranges", handler.CreateColorRanges)
	g.POST("/color-range", handler.CreateColorRange)
	g.PUT("/color-range", handler.UpdateColorRange)
	g.DELETE("/color-range", handler.DeleteColorRange)
}

// Categories
func (h *commonHandler) GetCategories(ctx echo.Context) error {
	var req model.CategoryRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.CategoryResponse
	if err := h.UUcase.GetCategories(ctx.Request().Context(), &req, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *commonHandler) CreateCategories(ctx echo.Context) error {
	var categories []*model.Category
	if err := ctx.Bind(&categories); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.CategoryResponse
	if err := h.UUcase.CreateCategories(ctx.Request().Context(), categories, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Categories)
}

func (h *commonHandler) CreateCategory(ctx echo.Context) error {
	var category model.Category
	if err := ctx.Bind(&category); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.CategoryResponse
	if err := h.UUcase.CreateCategory(ctx.Request().Context(), &category, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Category)
}

func (h *commonHandler) UpdateCategory(ctx echo.Context) error {
	var category model.Category
	if err := ctx.Bind(&category); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.CategoryResponse
	if err := h.UUcase.UpdateCategory(ctx.Request().Context(), &category, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Category)
}

func (h *commonHandler) DeleteCategory(ctx echo.Context) error {
	var category model.Category
	if err := ctx.Bind(&category); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.CategoryResponse
	if err := h.UUcase.DeleteCategory(ctx.Request().Context(), &category, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Category)
}

// Moods
func (h *commonHandler) GetMoods(ctx echo.Context) error {
	var req model.MoodRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.MoodResponse
	if err := h.UUcase.GetMoods(ctx.Request().Context(), &req, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *commonHandler) CreateMoods(ctx echo.Context) error {
	var moods []*model.Mood
	if err := ctx.Bind(&moods); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.MoodResponse
	if err := h.UUcase.CreateMoods(ctx.Request().Context(), moods, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Moods)
}

func (h *commonHandler) CreateMood(ctx echo.Context) error {
	var mood model.Mood
	if err := ctx.Bind(&mood); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.MoodResponse
	if err := h.UUcase.CreateMood(ctx.Request().Context(), &mood, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Mood)
}

func (h *commonHandler) UpdateMood(ctx echo.Context) error {
	var mood model.Mood
	if err := ctx.Bind(&mood); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.MoodResponse
	if err := h.UUcase.UpdateMood(ctx.Request().Context(), &mood, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Mood)
}

func (h *commonHandler) DeleteMood(ctx echo.Context) error {
	var mood model.Mood
	if err := ctx.Bind(&mood); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.MoodResponse
	if err := h.UUcase.DeleteMood(ctx.Request().Context(), &mood, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Mood)
}

// Color Ranges
func (h *commonHandler) GetColorRanges(ctx echo.Context) error {
	var req model.ColorRangeRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.ColorRangeResponse
	if err := h.UUcase.GetColorRanges(ctx.Request().Context(), &req, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *commonHandler) CreateColorRanges(ctx echo.Context) error {
	var colorRanges []*model.ColorRange
	if err := ctx.Bind(&colorRanges); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.ColorRangeResponse
	if err := h.UUcase.CreateColorRanges(ctx.Request().Context(), colorRanges, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.ColorRanges)
}

func (h *commonHandler) CreateColorRange(ctx echo.Context) error {
	var colorRange model.ColorRange
	if err := ctx.Bind(&colorRange); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.ColorRangeResponse
	if err := h.UUcase.CreateColorRange(ctx.Request().Context(), &colorRange, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.ColorRange)
}

func (h *commonHandler) UpdateColorRange(ctx echo.Context) error {
	var colorRange model.ColorRange
	if err := ctx.Bind(&colorRange); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.ColorRangeResponse
	if err := h.UUcase.UpdateColorRange(ctx.Request().Context(), &colorRange, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.ColorRange)
}

func (h *commonHandler) DeleteColorRange(ctx echo.Context) error {
	var colorRange model.ColorRange
	if err := ctx.Bind(&colorRange); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.ColorRangeResponse
	if err := h.UUcase.DeleteColorRange(ctx.Request().Context(), &colorRange, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.ColorRange)
}
