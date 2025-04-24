package handler

import (
	"github.com/amir-mirjalili/notes.git/service/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	user.Service
}

func NewUserHandler(s user.Service) *UserHandler {
	return &UserHandler{s}
}

func (h *UserHandler) List(c echo.Context) error {
	users, err := h.Service.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}
