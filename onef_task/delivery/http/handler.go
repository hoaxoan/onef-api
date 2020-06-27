package http

import (
	"net/http"

	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_task"

	"github.com/labstack/echo/v4"
)

type taskHandler struct {
	UUcase onef_task.Usecase
}

func NewHandler(e *echo.Echo, uc onef_task.Usecase) {
	handler := &taskHandler{
		UUcase: uc,
	}

	PublicRoute(e, handler)
}

func PublicRoute(e *echo.Echo, handler *taskHandler) {
	g := e.Group("/api/v1/task")
	g.PATCH("/create", handler.Create)
}

func (h *taskHandler) Create(ctx echo.Context) error {
	var task model.Task
	if err := ctx.Bind(&task); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.TaskResponse
	if err := h.UUcase.Create(ctx.Request().Context(), &task, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res)
}
