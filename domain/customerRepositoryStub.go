package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "Prosper Opara", City: "Portharcourt", Zipcode: "50000102", DOB: "04-02-1998", Status: "ACTIVE"},
		{Id: "2", Name: "Philemon Obi", City: "Portharcourt", Zipcode: "50000102", DOB: "04-02-1998", Status: "ACTIVE"},
	}

	return CustomerRepositoryStub{customers}
}
