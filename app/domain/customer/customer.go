package customer

import (
	"context"
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

// Handler provides customer related handler methods.
type Handler interface {
	Create(c *gin.Context)
}

// Service provides customer related service methods.
type Service interface {
	//
}

// Repository provides customer related repository methods.
type Repository interface {
	Create(ctx context.Context, c *Customer) (*Customer, error)
}
