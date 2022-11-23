package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Tiago", "São Paulo", "110011", "2000-01-01", "1"},
		{"1002", "João", "São Paulo", "110011", "2000-01-02", "1"},
	}

	return CustomerRepositoryStub{customers: customers}
}
