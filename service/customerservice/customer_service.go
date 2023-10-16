package customerservice

import (
	"github.com/kodekage/banking/dto"
	"github.com/kodekage/banking/internal/errors"
	"github.com/kodekage/banking/repositories/customerrepository"
)

type Service interface {
	GetAllCustomers(filters dto.CustomerFilters) ([]dto.CustomerResponse, *errors.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errors.AppError)
}

type customerService struct {
	repo customerrepository.Repository
}

var _ Service = (*customerService)(nil)

func New(repository customerrepository.Repository) Service {
	return customerService{repository}
}

func (s customerService) GetAllCustomers(filters dto.CustomerFilters) ([]dto.CustomerResponse, *errors.AppError) {
	customerFilter := ""

	if filters.Status == "active" {
		customerFilter = "1"
	}

	if filters.Status == "inactive" {
		customerFilter = "0"
	}

	data, err := s.repo.FindAll(customerFilter)
	if err != nil {
		return nil, err
	}

	result := make([]dto.CustomerResponse, 0)
	for i := 0; i < len(data); i++ {
		result = append(result, dto.CustomerResponse{
			Id:      data[i].Id,
			Name:    data[i].Name,
			City:    data[i].City,
			Zipcode: data[i].Zipcode,
			DOB:     data[i].DOB,
			Status:  data[i].Status,
		})
	}

	return result, nil
}

func (s customerService) GetCustomer(id string) (*dto.CustomerResponse, *errors.AppError) {
	data, err := s.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	result := dto.CustomerResponse{
		Id:      data.Id,
		Name:    data.Name,
		City:    data.City,
		Zipcode: data.Zipcode,
		DOB:     data.DOB,
		Status:  data.Status,
	}

	return &result, nil
}
