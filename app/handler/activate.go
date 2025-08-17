package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) ActivateBotHandler(ctx echo.Context) error {
	body, err := h.GetRequestBody(ctx)
	if err != nil {
		return err
	}
	path := h.GetPath(ctx)

	err = h.enqueue.Enqueue(path, body)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, Response{Message: "success"})
}
