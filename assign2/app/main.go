package main

import (
	"fmt"

	"github.com/somesh2905/pkg/assign2/domain"
	"github.com/somesh2905/pkg/assign2/mapstore"
)

type CustomerController struct {
	store domain.CustomerStore
}

func (cc CustomerController) Add(c domain.Customer) {
	err := cc.store.Create(c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New Customer has been created")
}

func (cc CustomerController) Update(id string, c domain.Customer) {
	err := cc.store.Update(id, c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Customer has been updated")
}

func (cc CustomerController) GetByID(id string) {
	c, err := cc.store.GetByID(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(c)
}
func (cc CustomerController) GetAll() {
	c, err := cc.store.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, val := range c {
		fmt.Println(val)
	}
}
func (cc CustomerController) Delete(id string) {
	err := cc.store.Delete(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("")
}

func main() {
	controller := CustomerController{
		store: mapstore.NewMapStore(),
	}
	customer_1 := domain.Customer{
		ID:    "somesh123",
		Name:  "somesh",
		Email: "somesh@gmail.com",
	}
	customer_2 := domain.Customer{
		ID:    "sles123",
		Name:  "sles",
		Email: "sles@gmail.com",
	}
	controller.Add(customer_1)
	controller.Add(customer_2)
	customer_1.Email = "somesh123@gmial.com"
	controller.Update("somesh123", customer_1)
	controller.GetByID("somesh123")
	controller.GetAll()
	controller.Delete("sles123")
	controller.GetAll()

}
