package route

import (
	"tvfrontalgwservice/app/enqueue"
	"tvfrontalgwservice/app/handler"

	"github.com/labstack/echo/v4"
)

func BnBot(app *echo.Echo, enqueue enqueue.IEnqueue) {
	handler := handler.NewHandler(enqueue)
	app.POST("/timeframe-exe-interval", handler.NewTradeHandler)
	app.POST("/activate-bot", handler.ActivateBotHandler)
	app.POST("/deactivate-bot", handler.DeactivateBotHandler)

}
