package utils

import (
	"context"
	"github.com/07sima07/scrap-api/proxy"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/fetch"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"net/http"
	"strconv"
	"time"
)

func Render(url string, header http.Header, proxy *proxy.Proxy, timeout time.Duration) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// http.Header to map[string], exclude content length coz chromedp return error
	headers := make(map[string]interface{})
	for k, v := range header {
		headers[string(k)] = string(v[0])
	}
	delete(headers, "Content-Length")

	o := append(chromedp.DefaultExecAllocatorOptions[:],
		//... any options here
		chromedp.ProxyServer(proxy.Address+":"+strconv.Itoa(proxy.Port)),
	)

	cx, cancel := chromedp.NewExecAllocator(context.Background(), o...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(cx)
	defer cancel()

	if proxy.Login != "" && proxy.Password != "" {
		chromedp.ListenTarget(ctx, func(ev interface{}) {
			go func() {
				switch ev := ev.(type) {
				case *fetch.EventAuthRequired:
					c := chromedp.FromContext(ctx)
					execCtx := cdp.WithExecutor(ctx, c.Target)

					resp := &fetch.AuthChallengeResponse{
						Response: fetch.AuthChallengeResponseResponseProvideCredentials,
						Username: proxy.Login,
						Password: proxy.Password,
					}

					fetch.ContinueWithAuth(ev.RequestID, resp).Do(execCtx)

				case *fetch.EventRequestPaused:
					c := chromedp.FromContext(ctx)
					execCtx := cdp.WithExecutor(ctx, c.Target)
					fetch.ContinueRequest(ev.RequestID).Do(execCtx)
				}
			}()
		})
	}

	var res string
	err := chromedp.Run(ctx,
		fetch.Enable().WithHandleAuthRequests(true),
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers(headers)),
		RunWithTimeOut(&ctx, timeout, chromedp.Tasks{
			chromedp.Navigate(url),
		}),
		chromedp.WaitReady(`body`),
		chromedp.OuterHTML("html", &res),
	)

	return res, err
}

func RunWithTimeOut(ctx *context.Context, timeout time.Duration, tasks chromedp.Tasks) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		timeoutContext, cancel := context.WithTimeout(ctx, timeout*time.Second)
		defer cancel()
		return tasks.Do(timeoutContext)
	}
}
