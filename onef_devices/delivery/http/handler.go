package http

import (
	"net/http"

	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_devices"

	"github.com/labstack/echo/v4"
)

type handler struct {
	UUcase onef_devices.Usecase
}

func NewHandler(e *echo.Echo, uc onef_devices.Usecase) {
	handler := &handler{
		UUcase: uc,
	}

	PublicRoute(e, handler)
}

func PublicRoute(e *echo.Echo, handler *handler) {
	g := e.Group("/api/v1/devices")

	g.GET("", handler.Get)
	g.GET("/:device_uuid", handler.GetDeviceWithUuid)
	g.POST("", handler.Create)
	g.PUT("", handler.Update)
	g.DELETE("", handler.Delete)
}

func (h *handler) Get(ctx echo.Context) error {
	var req model.DeviceRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.DeviceResponse
	if err := h.UUcase.Get(ctx.Request().Context(), &req, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.Devices)
}

func (h *handler) GetDeviceWithUuid(ctx echo.Context) error {
	deviceUuid := ctx.Param("device_uuid")

	var res model.DeviceResponse
	if err := h.UUcase.GetDeviceWithUuid(ctx.Request().Context(), deviceUuid, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.Device)
}

func (h *handler) Create(ctx echo.Context) error {
	var device model.Device
	if err := ctx.Bind(&device); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.DeviceResponse
	if err := h.UUcase.Create(ctx.Request().Context(), &device, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Device)
}

func (h *handler) Update(ctx echo.Context) error {
	var device model.Device
	if err := ctx.Bind(&device); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.DeviceResponse
	if err := h.UUcase.Update(ctx.Request().Context(), &device, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Device)
}

func (h *handler) Delete(ctx echo.Context) error {
	var device model.Device
	if err := ctx.Bind(&device); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.DeviceResponse
	if err := h.UUcase.Delete(ctx.Request().Context(), &device, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Device)
}
