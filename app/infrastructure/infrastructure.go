package infrastructure

import (
	"context"
	serviceconfig "tvfrontalgwservice/config"
)

type IInfrastructure interface {
	NewOrder(ctx context.Context, body []byte)
}

type infrastructure struct {
	bnBotFtConfig *serviceconfig.BnBotFt
}

func NewInfrastructure(bnBotFtConfig *serviceconfig.BnBotFt) IInfrastructure {
	return &infrastructure{
		bnBotFtConfig: bnBotFtConfig,
	}
}
