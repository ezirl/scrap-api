package main

import (
	"github.com/07sima07/mobile-proxy-api/sqldb"
	"github.com/07sima07/mobile-proxy-api/user"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	db := sqldb.ConnectDB()
	userRepo := user.NewRepo(db)

	h := user.NewBaseHandler(userRepo)

	router := httprouter.New()
	router.GET("/user/:id", h.User)
	router.GET("/", h.Scrap)

	log.Fatal(http.ListenAndServe(":8080", router))
}
