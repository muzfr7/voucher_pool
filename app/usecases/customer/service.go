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
