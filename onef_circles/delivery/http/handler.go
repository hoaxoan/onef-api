package http

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/hoaxoan/onef-api/onef_circles"
	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_core/setting"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type handler struct {
	UUcase onef_circles.Usecase
}

func NewHandler(e *echo.Echo, uc onef_circles.Usecase) {
	handler := &handler{
		UUcase: uc,
	}

	PublicRoute(e, handler)
}

func PublicRoute(e *echo.Echo, handler *handler) {
	JWTConfig := middleware.JWTConfig{
		SigningKey: []byte(setting.Config.JWTSecret.JWTKey),
		Claims:     &model.CustomClaims{},
	}

	g := e.Group("/api/v1/circles")
	g.Use(middleware.JWTWithConfig(JWTConfig))

	g.GET("", handler.Get)
	g.GET("/:circleId", handler.GetCircleWithId)
	g.POST("", handler.Create)
	g.PUT("/:circleId", handler.UpdateCircleWithId)
	g.DELETE("/:circleId", handler.DeleteCircleWithId)
	g.POST("/name-check", handler.CheckNameIsAvailable)
}

func (h *handler) Get(ctx echo.Context) error {
	var req model.CircleRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.CircleResponse
	if err := h.UUcase.Get(ctx.Request().Context(), &req, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.Circles)
}

func (h *handler) GetCircleWithId(ctx echo.Context) error {
	circleId, err := strconv.ParseInt(ctx.Param("circleId"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, model.Error{Code: http.StatusNotFound, Description: err.Error()})
	}

	var res model.CircleResponse
	if err := h.UUcase.GetCircleWithId(ctx.Request().Context(), circleId, &res); err != nil {
		return ctx.JSON(http.StatusNotFound, model.Error{Code: http.StatusNotFound, Description: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res.Circle)
}

func (h *handler) Create(ctx echo.Context) error {
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

	var circle model.Circle
	if err := ctx.Bind(&circle); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	circle.Creator = user
	circle.CreatorId = &user.Id

	var res model.CircleResponse
	if err := h.UUcase.Create(ctx.Request().Context(), &circle, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Circle)
}

func (h *handler) UpdateCircleWithId(ctx echo.Context) error {
	circleId, err := strconv.ParseInt(ctx.Param("circleId"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, model.Error{Code: http.StatusNotFound, Description: err.Error()})
	}

	var req model.UpdateCircleRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	req.Id = circleId

	var res model.CircleResponse
	if err := h.UUcase.UpdateCircleWithId(ctx.Request().Context(), &req, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Circle)
}

func (h *handler) DeleteCircleWithId(ctx echo.Context) error {
	circleId, err := strconv.ParseInt(ctx.Param("circleId"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, model.Error{Code: http.StatusNotFound, Description: err.Error()})
	}

	var res model.CircleResponse
	if err := h.UUcase.DeleteCircleWithId(ctx.Request().Context(), circleId, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res.Circle)
}

func (h *handler) CheckNameIsAvailable(ctx echo.Context) error {
	var circle model.Circle
	if err := ctx.Bind(&circle); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.CircleResponse
	if err := h.UUcase.CheckNameIsAvailable(ctx.Request().Context(), &circle, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	if res.Circle == nil {
		return ctx.JSON(http.StatusBadRequest, res.Circle)
	}

	return ctx.JSON(http.StatusOK, res.Circle)
}
