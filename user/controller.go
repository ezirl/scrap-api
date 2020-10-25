package user

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BaseHandler struct {
	userRepo Repo
}

func NewBaseHandler(userRepo Repo) *BaseHandler {
	return &BaseHandler{
		userRepo: userRepo,
	}
}

func (b *BaseHandler) User(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))
	user, _ := b.userRepo.FindByID(id)
	res, _ := json.Marshal(&user)
	println(user.CheckLimitBan())
	w.Write(res)
}
