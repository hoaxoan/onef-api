package http

import (
	"net/http"

	"github.com/hoaxoan/onef-api/onef_auth"
	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_core/setting"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type userHandler struct {
	UUcase onef_auth.Usecase
}

func NewHandler(e *echo.Echo, uc onef_auth.Usecase) {
	handler := &userHandler{
		UUcase: uc,
	}
	PrivateRoute(e, handler)

	PublicRoute(e, handler)
}

func PrivateRoute(e *echo.Echo, handler *userHandler) {
	JWTConfig := middleware.JWTConfig{
		SigningKey: []byte(setting.Config.JWTSecret.JWTKey),
		Claims:     &model.CustomClaims{},
	}

	g := e.Group("/api/v1/auth")
	g.Use(middleware.JWTWithConfig(JWTConfig))
	g.PUT("/user/update", handler.Update)
}

func PublicRoute(e *echo.Echo, handler *userHandler) {
	g := e.Group("/api/v1/auth")
	g.PATCH("/login", handler.Login)
	g.POST("/register", handler.Register)
}

func (h *userHandler) Login(ctx echo.Context) error {
	var user model.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var token model.Token
	err := h.UUcase.Auth(ctx.Request().Context(), &user, &token)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, token)
}

func (h *userHandler) Register(ctx echo.Context) error {
	var user model.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var res model.UserResponse
	if err := h.UUcase.Create(ctx.Request().Context(), &user, &res); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	var token model.Token
	err := h.UUcase.GetToken(ctx.Request().Context(), &user, &token)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: err.Error()})
	}

	return ctx.JSON(http.StatusOK, token)
}

func (h *userHandler) Update(ctx echo.Context) error {
	userToken := ctx.Get("user").(*jwt.Token)
	if userToken == nil || userToken.Valid == false {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "token invalid"})
	}
	claims := userToken.Claims.(*model.CustomClaims)

	var resToken model.Token
	reqToken := model.Token{Token: userToken.Raw}
	if err := h.UUcase.ValidateToken(ctx.Request().Context(), &reqToken, &resToken); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	var user model.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	user.Email = claims.User.Email

	var res model.UserResponse
	err := h.UUcase.Update(ctx.Request().Context(), &user, &res)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	return ctx.JSON(http.StatusOK, res.User)
}
