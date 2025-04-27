package handler

import (
	"github.com/amir-mirjalili/notes.git/notes/params"
	"github.com/amir-mirjalili/notes.git/notes/service"
	"github.com/labstack/echo/v4"
)

type NoteHandler struct {
	service.Service
}

func NewNoteHandler(s *service.Service) *NoteHandler {
	return &NoteHandler{*s}
}

func (h *NoteHandler) Create(c echo.Context) error {
	var req params.AddNotRequest
	if err := c.Bind(&req); err != nil {
	}
	return h.Service.Create(req)
}
