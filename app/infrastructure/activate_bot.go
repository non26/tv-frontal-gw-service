package infrastructure

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

func (i *infrastructure) ActivateBot(ctx context.Context, body []byte) {
	baseUrl := i.bnBotFtConfig.BaseUrl
	activateBotEndpoint := i.bnBotFtConfig.ActivateBotEndpoint

	fullUrl := fmt.Sprintf("%s%s", baseUrl, activateBotEndpoint)
	req, _ := http.NewRequest("POST", fullUrl, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	http.DefaultClient.Do(req)
}
