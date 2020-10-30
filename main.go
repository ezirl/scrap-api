package main

import (
	"github.com/07sima07/scrap-api/proxy"
	"github.com/07sima07/scrap-api/scrap"
	"github.com/07sima07/scrap-api/sqldb"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	db := sqldb.ConnectDB()

	// repositories
	proxyRepo := proxy.NewRepo(db)

	// controllers
	scrapHandler := scrap.NewBaseHandler(proxyRepo)

	// routers
	router := httprouter.New()
	router.GET("/", scrapHandler.Scrap)
	router.POST("/", scrapHandler.Scrap)
	router.PUT("/", scrapHandler.Scrap)
	router.PATCH("/", scrapHandler.Scrap)
	router.DELETE("/", scrapHandler.Scrap)
	router.OPTIONS("/", scrapHandler.Scrap)
	router.HEAD("/", scrapHandler.Scrap)

	log.Fatal(http.ListenAndServe(":8090", router))
}
