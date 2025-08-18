package processor

import "context"

func (p *processor) tvActivation(ctx context.Context, body []byte) {
	p.infrastructure.TvActivation(ctx, body)
}
