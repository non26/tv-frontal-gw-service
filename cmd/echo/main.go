package main

import (
	"fmt"
	"log"
	"tvfrontalgwservice/app/enqueue"
	"tvfrontalgwservice/app/handler"
	"tvfrontalgwservice/app/infrastructure"
	"tvfrontalgwservice/app/processor"
	serviceconfig "tvfrontalgwservice/config"

	"github.com/labstack/echo/v4"
)

func main() {
	config, err := serviceconfig.ReadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)

	}
	echo_app := echo.New()
	infrastructure := infrastructure.NewInfrastructure(&config.BnBotFt)
	processor := processor.NewProcessor(&config.BnBotFt, infrastructure)
	enqueue := enqueue.NewEnqueue(config.Queue.Size, processor)

	go enqueue.Do()

	handler := handler.NewHandler(enqueue)
	echo_app.POST("/timeframe-exe-interval", handler.NewTradeHandler)

	echo_app.Logger.Fatal(echo_app.Start(fmt.Sprintf(":%d", config.Server.Port)))
	defer enqueue.Close()
}
