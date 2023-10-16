package entities

type Customer struct {
	Id      string `db:"customer_id"`
	Name    string
	City    string
	Zipcode string
	DOB     string `db:"date_of_birth"`
	Status  string
}
