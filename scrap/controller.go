package scrap

import (
	"compress/flate"
	"compress/gzip"
	"github.com/07sima07/scrap-api/proxy"
	"github.com/07sima07/scrap-api/utils"
	"github.com/andybalholm/brotli"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type BaseHandler struct {
	proxyRepo proxy.Repo
}

func NewBaseHandler(proxyRepo proxy.Repo) *BaseHandler {
	return &BaseHandler{
		proxyRepo: proxyRepo,
	}
}

func (b *BaseHandler) Scrap(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	params := r.URL.Query()
	clientUrl := params["url"][0]

	if !strings.Contains(clientUrl, "http") {
		clientUrl = "http://" + clientUrl
	}

	if params["render"] != nil {
		res, err := utils.Render(clientUrl, r.Header.Clone())
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"status": "error", "msg": "wrong url or url dont response"}`))
		}

		w.Write([]byte(res))
		return
	}

	req, _ := http.NewRequest(r.Method, clientUrl, nil)
	req.Header = r.Header.Clone()
	req.Body = r.Body

	httpProxy, err := b.getProxy(params)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"status": "error", "msg": "proxy not found"}`))
		return
	}

	proxyUrl, _ := url.Parse("http://"+httpProxy.Address)
	client := http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
		Timeout: 60 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"status": "error", "msg": "wrong url or url dont response"}`))
		return
	}
	defer res.Body.Close()

	body := encodingResult(res)
	w.Write(body)
}

func (b *BaseHandler) getProxy(params url.Values) (*proxy.Proxy, error) {
	if params["country"] != nil {
		return b.proxyRepo.FindByCountry(params["country"][0])
	}
	return b.proxyRepo.GetRandom()
}

// decode response
func encodingResult(res *http.Response) []byte {
	var reader io.Reader
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(res.Body)
	case "br":
		reader = brotli.NewReader(res.Body)
	case "deflate":
		reader = flate.NewReader(res.Body)
	default:
		reader = res.Body
	}

	body, _ := ioutil.ReadAll(reader)
	return body
}
