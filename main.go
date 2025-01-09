package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router = routesInit(router)

	err := http.ListenAndServe(":8080", router)
	if err == nil {
		fmt.Println("server running on port 8080...")
	}
}
