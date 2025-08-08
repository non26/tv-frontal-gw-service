package processor

import (
	"context"
	"tvfrontalgwservice/app/infrastructure"
	serviceconfig "tvfrontalgwservice/config"
)

type IProcessor interface {
	ProcessTask(path string, body []byte) error
	newOrder(ctx context.Context, body []byte)
}

type processor struct {
	bnBotFtConfig  *serviceconfig.BnBotFt
	infrastructure infrastructure.IInfrastructure
}

func NewProcessor(
	bnBotFtConfig *serviceconfig.BnBotFt,
	infrastructure infrastructure.IInfrastructure,
) IProcessor {
	return &processor{
		bnBotFtConfig:  bnBotFtConfig,
		infrastructure: infrastructure,
	}
}

func (p *processor) ProcessTask(path string, body []byte) error {
	switch path {
	case p.bnBotFtConfig.NewOrderEndpoint:
		p.infrastructure.NewOrder(context.Background(), body)
	}
	return nil
}
