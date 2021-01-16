package main

import (
	"github.com/07sima07/scrap-api/proxy"
	"github.com/07sima07/scrap-api/scrap"
	"github.com/07sima07/scrap-api/sqldb"
	"github.com/07sima07/scrap-api/user"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var authUser = "kosmos"
var authPassword = "AbubuAeAkak322"

func main() {
	db := sqldb.ConnectDB()

	// repositories
	proxyRepo := proxy.NewRepo(db)
	userRepo := user.NewRepo(db)
	callRepo := scrap.NewRepo(db, userRepo)

	// controllers
	scrapHandler := scrap.NewBaseHandler(proxyRepo, callRepo, userRepo)
	proxyHandler := proxy.NewBaseHandler(proxyRepo)
	userHandler := user.NewBaseHandler(userRepo)

	// routers
	router := httprouter.New()

	router.GET("/proxies", BasicAuth(proxyHandler.GetProxies, authUser, authPassword))
	router.POST("/proxy/create", BasicAuth(proxyHandler.CreateProxy, authUser, authPassword))
	router.GET("/proxy/:id/delete", BasicAuth(proxyHandler.DeleteProxy, authUser, authPassword))
	router.GET("/calls", BasicAuth(scrapHandler.GetCalls, authUser, authPassword))
	router.GET("/users", BasicAuth(userHandler.GetUsers, authUser, authPassword))

	router.GET("/", scrapHandler.Scrap)
	router.POST("/", scrapHandler.Scrap)
	router.PUT("/", scrapHandler.Scrap)
	router.PATCH("/", scrapHandler.Scrap)
	router.DELETE("/", scrapHandler.Scrap)
	router.OPTIONS("/", scrapHandler.Scrap)
	router.HEAD("/", scrapHandler.Scrap)

	log.Fatal(http.ListenAndServe(":8090", router))
}
