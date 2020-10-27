package scrap

import (
	"github.com/07sima07/mobile-proxy-api/user"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BaseHandler struct {
	userRepo user.Repo
}

func NewBaseHandler(userRepo user.Repo) *BaseHandler {
	return &BaseHandler{
		userRepo: userRepo,
	}
}

func (b *BaseHandler) Scrap(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}
