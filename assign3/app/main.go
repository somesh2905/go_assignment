package main

import (
	"log"
	"net/http"

	"github.com/somesh2905/pkg/assign3/router"
)

func main() {

	log.Println("Listening !")
	router.Routers()
	http.ListenAndServe(":8080", router.Mux)

}
