package handler

import (
	"io"
	"tvfrontalgwservice/app/enqueue"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string `json:"message"`
}

type IHandler interface {
	GetRequestBody(ctx echo.Context) ([]byte, error)
	GetPath(ctx echo.Context) string
	NewTradeHandler(ctx echo.Context) error
	ActivateBotHandler(ctx echo.Context) error
	DeactivateBotHandler(ctx echo.Context) error
	TvActivationHandler(ctx echo.Context) error
}

type handler struct {
	enqueue enqueue.IEnqueue
}

func NewHandler(enqueue enqueue.IEnqueue) IHandler {
	return &handler{
		enqueue: enqueue,
	}
}

func (h *handler) GetRequestBody(ctx echo.Context) ([]byte, error) {
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}
	ctx.Request().Body.Close()
	return body, nil
}

func (h *handler) GetPath(ctx echo.Context) string {
	return ctx.Request().URL.Path
}
