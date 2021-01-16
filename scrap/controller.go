package scrap

import (
	"compress/flate"
	"compress/gzip"
	"encoding/json"
	"github.com/07sima07/scrap-api/proxy"
	"github.com/07sima07/scrap-api/user"
	"github.com/07sima07/scrap-api/utils"
	"github.com/andybalholm/brotli"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type BaseHandler struct {
	proxyRepo proxy.Repo
	userRepo  user.Repo
	callRepo  Repo
}

func NewBaseHandler(proxyRepo proxy.Repo, callRepo Repo, userRepo user.Repo) *BaseHandler {
	return &BaseHandler{
		proxyRepo: proxyRepo,
		callRepo:  callRepo,
		userRepo:  userRepo,
	}
}

func (b *BaseHandler) Scrap(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	params := r.URL.Query()
	if _, ok := params["url"]; !ok {
		w.Write([]byte(`{"status": "ok", "msg": "api work"}`))
		return
	}

	// check token
	if params["token"] == nil {
		w.Write([]byte(`{"status": "error", "msg": "wrong token"}`))
		return
	}

	// find user by token
	userModel, err := b.userRepo.FindByToken(params["token"][0])
	if err != nil || userModel.ID == 0 {
		w.Write([]byte(`{"status": "error", "msg": "wrong token"}`))
		return
	}

	clientUrl := params["url"][0]

	// check http protocol
	if !strings.Contains(clientUrl, "http") {
		clientUrl = "http://" + clientUrl
	}

	// check timeout
	var timeout int
	if params["timeout"] == nil {
		timeout = 60
	} else {
		tint, err := strconv.Atoi(params["timeout"][0])
		timeout = tint
		if err != nil || params["timeout"][0] == "0" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"status": "error", "msg": "wrong timeout value"}`))
			return
		}
	}


	// check proxies
	httpProxy, err := b.getProxy(params)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"status": "error", "msg": "proxy not found"}`))
		return
	}

	proxyStr := getProxyStr(httpProxy)

	// check render param and send response if render on
	proxyUrl, _ := url.Parse(proxyStr)
	if params["render"] != nil {
		res, err := utils.Render(clientUrl, r.Header.Clone(), httpProxy, time.Duration(timeout))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"status": "error", "msg": "wrong url or url dont response"}`))
		}

		// increment user requests and save api call
		b.userRepo.IncRequests(userModel.ID)
		urlCall := clientUrl
		if len(urlCall) > 4999 {
			urlCall = urlCall[0:5000]
		}
		call := Call{
			UserID: userModel.ID,
			Url:    urlCall,
		}
		err = b.callRepo.Save(call)

		if err != nil {
			log.Println("Api call is not saved")
		}

		w.Write([]byte(res))
		return
	}


	// send response without render
	req, _ := http.NewRequest(r.Method, clientUrl, nil)
	req.Header = r.Header.Clone()
	req.Body = r.Body

	client := http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
		Timeout:   time.Duration(timeout) * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"status": "error", "msg": "wrong url or url dont response"}`))
		return
	}
	defer res.Body.Close()

	// increment user requests and save api call
	b.userRepo.IncRequests(userModel.ID)
	urlCall := clientUrl
	if len(urlCall) > 4999 {
		urlCall = urlCall[0:5000]
	}
	call := Call{
		UserID: userModel.ID,
		Url:    urlCall,
	}
	err = b.callRepo.Save(call)

	if err != nil {
		log.Println("Api call is not saved")
	}

	body := encodingResult(res)
	w.Write(body)
}

func getProxyStr(httpProxy *proxy.Proxy) string {
	proxyStr := httpProxy.Type + "://"
	if httpProxy.Login != "" && httpProxy.Password != "" {
		proxyStr += httpProxy.Login + ":" + httpProxy.Password + "@"
	}
	proxyStr += httpProxy.Address + ":" + httpProxy.Port
	return proxyStr
}

func (b *BaseHandler) getProxy(params url.Values) (*proxy.Proxy, error) {
	if params["country"] != nil {
		return b.proxyRepo.FindByCountry(params["country"][0])
	}
	return b.proxyRepo.GetRandom()
}

func (b *BaseHandler) GetCalls(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	calls, _ := b.callRepo.All(10)
	res, _ := json.Marshal(&calls)
	writer.Write(res)
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
