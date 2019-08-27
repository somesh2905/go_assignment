package router

import (
	"github.com/gorilla/mux"
	"github.com/somesh2905/pkg/assign3/controller"
)

var Mux = mux.NewRouter()

func Routers() {
	Mux.HandleFunc("/Create", controller.Create)
	Mux.HandleFunc("/Update", controller.Update)
	Mux.HandleFunc("/GetById", controller.GetByID)
	Mux.HandleFunc("/GetAll", controller.GetAll)
	Mux.HandleFunc("/Delete", controller.Delete)
}
