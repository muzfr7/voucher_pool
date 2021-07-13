package repository

import (
	"context"

	customerDomain "github.com/muzfr7/voucher_pool/app/domain/customer"
)

// repository implements customerDomain.Repository interface.
type repository struct{}

// NewService returns a new instance of service.
func NewCustomerRepository() customerDomain.Repository {
	return &repository{}
}

// Create will store a customer in database.
func (r *repository) Create(ctx context.Context, c *customerDomain.Customer) (*customerDomain.Customer, error) {
	return nil, nil
}
