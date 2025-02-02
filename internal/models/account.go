package models

import "gorm.io/gorm"

type Account struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	Balance float64 `gorm:"type:decimal(10,2);not null;default:0" json:"balance"`
}

// Migrate the account table
func MigrateAccount(db *gorm.DB) error {
	return db.AutoMigrate(&Account{})
}
