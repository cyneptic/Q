package controller

import (
	"net/http"
	"time"

	"github.com/cyneptic/letsgo-authentication/controller/middleware"
	"github.com/cyneptic/letsgo-authentication/controller/validators"
	"github.com/cyneptic/letsgo-authentication/internal/core/entities"
	"github.com/cyneptic/letsgo-authentication/internal/core/ports"
	"github.com/cyneptic/letsgo-authentication/internal/core/service"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

type DisableUserRequest struct {
	TargetID string `json:"target_id"`
	Toggle   bool   `json:"toggle"`
}

type AuthenticationHandler struct {
	svc ports.UserServiceContract
}

func NewAuthenticationHandler() *AuthenticationHandler {
	svc := service.NewAuthenticationService()
	return &AuthenticationHandler{
		svc: svc,
	}
}
func AddAuthServiceRoutes(e *echo.Echo) {
	h := NewAuthenticationHandler()
	e.POST("/login", h.loginHandler)
	e.POST("/register", h.register)
	e.POST("/logout", h.logout)
	e.POST("/create-admin", h.CreateAdmin)
	e.GET("/is-admin/:id/:role", h.IsAdmin)
	e.GET("/verify/:number/:id", h.Verify)
	e.POST("/disable-user", h.DisableUser)
	e.GET("/test", h.Test, middleware.AuthMiddleware)
}
func (h *AuthenticationHandler) Test(c echo.Context) error {
	return c.JSON(http.StatusOK, "You Hit /test so token is valid")
}

func (h *AuthenticationHandler) DisableUser(c echo.Context) error {
	var request DisableUserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := validators.DisableUser(request.TargetID, request.Toggle)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	target, err := uuid.Parse(request.TargetID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = h.svc.DisableUser(target, request.Toggle)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, nil)
}

// validation done
func (h *AuthenticationHandler) loginHandler(c echo.Context) error {
	user := new(entities.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	err := validators.LoginValidation(*user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	token, id, err := h.svc.LoginService(*user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Error": "Invalid Email Or Password",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token, "id": id})
}

// validation done
func (h *AuthenticationHandler) logout(c echo.Context) error {

	authHeader := c.Request().Header.Get("Authorization")

	err := validators.LogoutValidation(authHeader)

	if err != nil {
		return c.JSON(http.StatusForbidden, err.Error())
	}

	err = h.svc.Logout(authHeader)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "logout successful")
}

// validation done
func (h *AuthenticationHandler) register(c echo.Context) error {

	newUser := &entities.User{
		DBModel: entities.DBModel{
			CreatedAt: time.Now(),
		},
		Role: "user",
	}

	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	err := validators.RegisterValidation(*newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.svc.AddUser(*newUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]*entities.User{
		"newUser": newUser,
	})
}

// validation done
func (h *AuthenticationHandler) CreateAdmin(c echo.Context) error {
	newAdmin := &entities.User{
		DBModel: entities.DBModel{
			CreatedAt: time.Now(),
		},
	}

	if err := c.Bind(&newAdmin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}
	err := validators.RegisterValidation(*newAdmin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = h.svc.AddUser(*newAdmin)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, newAdmin)
}

// validation done
func (h *AuthenticationHandler) IsAdmin(c echo.Context) error {
	idParams := c.Param("id")
	role := c.Param("role")
	err := validators.IsAdmin(idParams, role)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	accountId, err := uuid.Parse(idParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	isAdmin, err := h.svc.IsAdminAccount(accountId, role)
	if err != nil {
		println(err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, isAdmin)
}

// validation done
func (h *AuthenticationHandler) Verify(c echo.Context) error {
	number := c.Param("number")
	id := c.Param("id")

	err := validators.VerifyValidation(number, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	accountId, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	verifiedAccount, err := h.svc.Verify(number, accountId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, verifiedAccount)
}
