package service

import (
	"github.com/kodekage/banking/domain"
	"github.com/kodekage/banking/dto"
	"github.com/kodekage/banking/errs"
)

type CustomerService interface {
	GetAllCustomers(filters domain.Filters) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(filters domain.Filters) ([]dto.CustomerResponse, *errs.AppError) {
	if filters.Status == "active" {
		filters.Status = "1"
	}

	data, err := s.repo.FindAll(filters)
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

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
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

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
