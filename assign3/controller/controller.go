package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/somesh2905/pkg/assign3/domain"
	"github.com/somesh2905/pkg/assign3/mapstore"
)

type CustomerController struct {
	store domain.CustomerStore
}

var controller = CustomerController{
	store: mapstore.NewMapStore(),
}

func Create(w http.ResponseWriter, r *http.Request) {
	c := domain.Customer{
		ID:    "somesh123",
		Name:  "somesh",
		Email: "somesh@gmail.com",
	}
	err := controller.store.Create(c)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("New Customer has been created")
	rs, _ := json.Marshal(c)
	w.Write(rs)
}

func Update(w http.ResponseWriter, r *http.Request) {
	c := domain.Customer{
		ID:    "somesh123",
		Name:  "somesh",
		Email: "somesh56@gmail.com",
	}
	err := controller.store.Update(c.Email, c)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Customer has been updated")
	rs, _ := json.Marshal("user updated successfully")
	w.Write(rs)

}

func GetByID(w http.ResponseWriter, r *http.Request) {
	c := domain.Customer{
		ID: "somesh123",
	}
	c, err := controller.store.GetByID(c.ID)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	rs, _ := json.Marshal(c)
	w.Write(rs)

	fmt.Println(c)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	c, err := controller.store.GetAll()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, val := range c {

		fmt.Println(val)
	}
	rs, _ := json.Marshal(c)
	w.Write(rs)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	c := domain.Customer{
		ID: "somesh123",
	}
	err := controller.store.Delete(c.ID)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(c)
	rs, _ := json.Marshal("deleted successfully")
	w.Write(rs)

}
