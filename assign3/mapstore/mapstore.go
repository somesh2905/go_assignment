package mapstore

import (
	"errors"
	"fmt"

	"github.com/somesh2905/pkg/assign3/domain"
)

// MapStore is a simple structure of map type
type MapStore struct {
	Store map[string]domain.Customer
}

//NewMapStore is used to create mapstore
func NewMapStore() *MapStore {
	return &MapStore{Store: make(map[string]domain.Customer)}
}

//Create is used to enter customer info in store
func (ms *MapStore) Create(c domain.Customer) error {
	if c.ID == "" {
		return errors.New("error")
	}
	ms.Store[c.ID] = c
	return nil

}

//Update is used to update customer info
func (ms *MapStore) Update(id string, c domain.Customer) error {
	if id == "" {
		return errors.New("error")
	}
	ms.Store[id] = c
	fmt.Println("updated id", id)
	return nil
}

//GetByID is used to retrieve customer info by ID
func (ms *MapStore) GetByID(id string) (domain.Customer, error) {
	if id == "" {

		return domain.Customer{}, errors.New("id not found")
	}

	c, ok := ms.Store[id]
	if !ok {
		return domain.Customer{}, errors.New("id not found")
	}
	return c, nil
}

//GetAll is used to retrieve all the customer info
func (ms *MapStore) GetAll() ([]domain.Customer, error) {
	if len(ms.Store) == 0 {
		return []domain.Customer{}, errors.New("empty store")
	}
	clist := []domain.Customer{}
	for _, c := range ms.Store {
		clist = append(clist, c)

	}
	return clist, nil
}

//Delete is used to delete a customer entry by ID
func (ms *MapStore) Delete(id string) error {
	if id == "" {
		return errors.New("error")
	}
	delete(ms.Store, id)
	return nil
}
