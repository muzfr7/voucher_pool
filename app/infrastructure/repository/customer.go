package repository

import (
	"github.com/jinzhu/gorm"
	customerDomain "github.com/muzfr7/voucher_pool/app/domain/customer"
)

// repository implements customerDomain.Repository interface.
type repository struct {
	db *gorm.DB
}

// NewService returns a new instance of service.
func NewCustomerRepository(db *gorm.DB) customerDomain.Repository {
	return &repository{
		db: db,
	}
}

// Create will store a customer in database.
func (r *repository) Create(c *customerDomain.Customer) error {
	err := r.db.Table(c.TableName()).Create(c).Error
	if err != nil {
		return err
	}

	return nil
}
