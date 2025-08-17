package infrastructure

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

func (i *infrastructure) DeactivateBot(ctx context.Context, body []byte) {
	baseUrl := i.bnBotFtConfig.BaseUrl
	deactivateBotEndpoint := i.bnBotFtConfig.DeactivateBotEndpoint

	fullUrl := fmt.Sprintf("%s%s", baseUrl, deactivateBotEndpoint)
	req, _ := http.NewRequest("POST", fullUrl, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	http.DefaultClient.Do(req)
}
