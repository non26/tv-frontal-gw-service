package processor

import "context"

func (p *processor) deactivateBot(ctx context.Context, body []byte) {
	p.infrastructure.DeactivateBot(ctx, body)
}
