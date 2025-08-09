package main

import (
	"context"
	"log"
	"tvfrontalgwservice/app/enqueue"
	"tvfrontalgwservice/app/infrastructure"
	"tvfrontalgwservice/app/processor"
	serviceconfig "tvfrontalgwservice/config"
	route "tvfrontalgwservice/route/bn"
	routehealthcheck "tvfrontalgwservice/route/health_check"
	routeupdateconfig "tvfrontalgwservice/route/update_config"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
)

var echoLambda *echoadapter.EchoLambda
var _enqueue enqueue.IEnqueue

func init() {
	config, err := serviceconfig.ReadAWSAppConfig()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)

	}
	echo_app := echo.New()
	infrastructure := infrastructure.NewInfrastructure(&config.BnBotFt)
	processor := processor.NewProcessor(&config.BnBotFt, infrastructure)
	_enqueue = enqueue.NewEnqueue(config.Queue.Size, processor)

	go _enqueue.Do()

	route.BnBot(echo_app, _enqueue)
	routeupdateconfig.UpdateAWSAppConfig(echo_app, config)
	routehealthcheck.HealthCheck(echo_app)

	echoLambda = echoadapter.New(echo_app)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
	defer _enqueue.Close()
}
