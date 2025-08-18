package infrastructure

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

func (i *infrastructure) TvActivation(ctx context.Context, body []byte) {
	baseUrl := i.bnBotFtConfig.BaseUrl
	tvActivationEndpoint := i.bnBotFtConfig.TvActivationEndpoint

	fullUrl := fmt.Sprintf("%s%s", baseUrl, tvActivationEndpoint)
	req, _ := http.NewRequest("POST", fullUrl, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	http.DefaultClient.Do(req)
}
