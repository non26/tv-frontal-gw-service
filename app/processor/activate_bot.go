package processor

import "context"

func (p *processor) activateBot(ctx context.Context, body []byte) {
	p.infrastructure.ActivateBot(ctx, body)
}
