package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	AccountID uint      `gorm:"not null" json:"account_id"`
	Type      string    `gorm:"type:varchar(10);not null;check:type IN ('deposit', 'withdraw')" json:"type"`
	Amount    float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
}

// Migrate the transactions table
func MigrateTransaction(db *gorm.DB) error {
	return db.AutoMigrate(&Transaction{})
}
