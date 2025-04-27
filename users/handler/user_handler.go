package handler

import (
	"github.com/amir-mirjalili/notes.git/users/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	service.Service
}

func NewUserHandler(s service.Service) *UserHandler {
	return &UserHandler{s}
}

func (h *UserHandler) List(c echo.Context) error {
	users, err := h.Service.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}
