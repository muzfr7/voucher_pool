package voucher

import "time"

// Voucher represents a voucher in database.
type Voucher struct {
	ID         uint      `json:"id"`
	CustomerID uint      `json:"customer_id" gorm:"default:null"`
	OfferID    uint      `json:"offer_id" gorm:"default:null"`
	Code       string    `json:"code"`
	UsedAt     time.Time `json:"used_at" gorm:"default:null"`
	ExpiresAt  time.Time `json:"expires_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// TableName return name of this table in database.
func (Voucher) TableName() string {
	return "vouchers"
}
