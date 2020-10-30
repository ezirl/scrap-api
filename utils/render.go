package utils

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"net/http"
)

func Render(url string, header http.Header) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// http.Header to map[string], exclude content length coz chromedp return error
	headers := make(map[string]interface{})
	for k, v := range header {
		headers[string(k)] = string(v[0])
	}
	delete(headers, "Content-Length")

	var res string
	err := chromedp.Run(ctx,
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers(headers)),
		chromedp.Navigate(url),
		chromedp.WaitReady(`body`),
		chromedp.OuterHTML("html", &res),
	)

	return res, err
}
