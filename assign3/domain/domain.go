package domain

// Customer contains details of customer
type Customer struct {
	ID    string
	Name  string
	Email string
}

// CustomerStore is an interface that does CRUD
// for Customer.
type CustomerStore interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	GetByID(string) (Customer, error)
	GetAll() ([]Customer, error)
}
