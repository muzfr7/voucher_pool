package customer

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Customer represents a customer in database.
type Customer struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName return name of this table in database.
func (Customer) TableName() string {
	return "customers"
}

// Handler defines contracts for customer handler.
type Handler interface {
	Create(c *gin.Context)
}

// Service defines contracts for customer service.
type Service interface {
	Create(payload *CreateCustomerRequestDTO) error
}

// Repository defines contracts for customer repository.
type Repository interface {
	Create(c *Customer) error
}
