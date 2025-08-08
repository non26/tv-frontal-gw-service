package infrastructure

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

func (i *infrastructure) NewOrder(ctx context.Context, body []byte) {
	baseUrl := i.bnBotFtConfig.BaseUrl
	newOrderEndpoint := i.bnBotFtConfig.NewOrderEndpoint

	fullUrl := fmt.Sprintf("%s%s", baseUrl, newOrderEndpoint)
	req, _ := http.NewRequest("POST", fullUrl, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	http.DefaultClient.Do(req)
}
