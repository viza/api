package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Vetal", City: "Kharkiv", Zipcode: "61029", Dateofbirth: "1982-02-23", Status: "1"},
		{Id: "1002", Name: "Rob", City: "Kyiv", Zipcode: "32000", Dateofbirth: "1985-05-25", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
