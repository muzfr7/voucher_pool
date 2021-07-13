package offer

import "time"

// Offer represents an offer in database.
type Offer struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Discount  uint      `json:"discount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName return name of this table in database.
func (Offer) TableName() string {
	return "offers"
}
