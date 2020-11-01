package proxy

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BaseHandler struct {
	proxyRepo Repo
}

func NewBaseHandler(userRepo Repo) *BaseHandler {
	return &BaseHandler{
		proxyRepo: userRepo,
	}
}

func (b *BaseHandler) CreateProxy(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	proxy := Proxy{}
	decoder.Decode(&proxy)
	fmt.Println(proxy)

	err := b.proxyRepo.Save(proxy)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"status": "error", "msg": "proxy not created"}`))
	}
	writer.Write([]byte(`{"status": "ok", "msg": "proxy created"}`))
}

func (b *BaseHandler) DeleteProxy(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"status": "error", "msg": "wrong id value"}`))
	}

	err = b.proxyRepo.Delete(id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"status": "error", "msg": "proxy not deleted"}`))
	}
	writer.WriteHeader(200)
	writer.Write([]byte(`{"status": "ok", "msg": "proxy deleted"}`))
}

func (b *BaseHandler) GetProxies(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	proxies, _ := b.proxyRepo.All()
	res, _ := json.Marshal(&proxies)
	writer.Write(res)
}