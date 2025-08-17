package main

import (
	"fmt"
	"log"
	"tvfrontalgwservice/app/enqueue"
	"tvfrontalgwservice/app/infrastructure"
	"tvfrontalgwservice/app/processor"
	serviceconfig "tvfrontalgwservice/config"
	route "tvfrontalgwservice/route/bn"
	routehealthcheck "tvfrontalgwservice/route/health_check"
	routeupdateconfig "tvfrontalgwservice/route/update_config"

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
	defer enqueue.Close()

	route.BnBot(echo_app, enqueue)
	routeupdateconfig.UpdateAWSAppConfig(echo_app, config)
	routehealthcheck.HealthCheck(echo_app)

	echo_app.Logger.Fatal(echo_app.Start(fmt.Sprintf(":%d", config.Server.Port)))
}
