package entities

import (
	"github.com/kodekage/banking/dto"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.CreateAccountResponse {
	return dto.CreateAccountResponse{AccountId: a.AccountId}
}
