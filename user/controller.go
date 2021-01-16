package user

import (
	"encoding/json"
	"fmt"
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

func (b *BaseHandler) GetUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	users, _ := b.userRepo.All(10)
	res, _ := json.Marshal(&users)
	w.Write(res)
}

func (b *BaseHandler) User(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))
	user, _ := b.userRepo.FindByID(id)
	res, _ := json.Marshal(&user)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"status": "error", "msg": "not found"}`))
	} else {
		w.Write(res)
	}
}

func (b *BaseHandler) CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	user := User{}

	decoder.Decode(&user)
	fmt.Println(user)
	b.userRepo.HashPassword(&user)

	b.userRepo.Save(user)
	writer.Write([]byte(`{"status": "ok", "msg": "user created"}`))
}
