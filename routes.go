package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/observer/pkg"
)

func routesInit(router *httprouter.Router) *httprouter.Router {
	router.GET("/query", pkg.GetData)
	router.POST("/ingest", pkg.StoreData)
	router.GET("/refresh", pkg.RefreshData)

	return router
}
