package processor

import "context"

func (p *processor) newOrder(ctx context.Context, body []byte) {
	p.infrastructure.NewOrder(ctx, body)
}
