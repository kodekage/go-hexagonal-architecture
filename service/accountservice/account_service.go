package accountservice

import (
	"time"

	"github.com/kodekage/banking/domain/entities"
	"github.com/kodekage/banking/dto"
	"github.com/kodekage/banking/internal/errors"
	"github.com/kodekage/banking/repositories/accountrepository"
)

type Service interface {
	CreateNewAccount(dto.CreateAccountRequest) (*dto.CreateAccountResponse, *errors.AppError)
}

type accountService struct {
	repo accountrepository.Repository
}

var _ Service = (*accountService)(nil)

func New(repository accountrepository.Repository) Service {
	return accountService{repository}
}

func (d accountService) CreateNewAccount(req dto.CreateAccountRequest) (*dto.CreateAccountResponse, *errors.AppError) {
	account := entities.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05Z07:00"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	newAccount, err := d.repo.Save(account)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDto()

	return &response, nil
}
