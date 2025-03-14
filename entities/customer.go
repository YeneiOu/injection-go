package entities

type Customer struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(id string) (Customer, error)
}
