package proxy

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
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
