package customer

import (
	customerDomain "github.com/muzfr7/voucher_pool/app/domain/customer"
)

// service implements customerDomain.Service interface.
type service struct {
	repo customerDomain.Repository
}

// NewService returns a new instance of service.
func NewService(r customerDomain.Repository) customerDomain.Service {
	return &service{
		repo: r,
	}
}

// Create will create a new customer.
func (s *service) Create(payload *customerDomain.CreateCustomerRequestDTO) error {
	// store customer in the database
	err := s.repo.Create(&customerDomain.Customer{
		Name:  payload.Name,
		Email: payload.Email,
	})
	if err != nil {
		return err
	}

	return nil
}
