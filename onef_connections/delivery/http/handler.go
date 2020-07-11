package http

import (
	"net/http"

	"github.com/hoaxoan/onef-api/onef_connections"
	"github.com/hoaxoan/onef-api/onef_core/model"

	"github.com/labstack/echo/v4"
)

type handler struct {
	UUcase onef_connections.Usecase
}

func NewHandler(e *echo.Echo, uc onef_connections.Usecase) {
	handler := &handler{
		UUcase: uc,
	}

	PublicRoute(e, handler)
}

func PublicRoute(e *echo.Echo, handler *handler) {
	g := e.Group("/api/v1/connections")
	// Categories
	g.GET("", handler.Get)
	g.POST("", handler.Create)
	g.PUT("", handler.Update)
	g.DELETE("", handler.Delete)
}

func (h *handler) Get(ctx echo.Context) error {
	var req model.ConnectionRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.ConnectionResponse
	if err := h.UUcase.Get(ctx.Request().Context(), &req, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.Connections)
}

func (h *handler) Create(ctx echo.Context) error {
	var connection model.Connection
	if err := ctx.Bind(&connection); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.ConnectionResponse
	if err := h.UUcase.Create(ctx.Request().Context(), &connection, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Connection)
}

func (h *handler) Update(ctx echo.Context) error {
	var connection model.Connection
	if err := ctx.Bind(&connection); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.ConnectionResponse
	if err := h.UUcase.Update(ctx.Request().Context(), &connection, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Connection)
}

func (h *handler) Delete(ctx echo.Context) error {
	var connection model.Connection
	if err := ctx.Bind(&connection); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.ConnectionResponse
	if err := h.UUcase.Delete(ctx.Request().Context(), &connection, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Connection)
}
