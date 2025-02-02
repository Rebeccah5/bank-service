package repository

import (
	"bank-service/internal/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type AccountRepository struct {
	DB *gorm.DB
}

// NewAccountRepository initializes the repository
func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{DB: db}
}

// GetBalance retrieves the current balance
func (r *AccountRepository) GetBalance() (float64, error) {
	var account models.Account
	err := r.DB.First(&account).Error
	if err != nil {
		return 0, err
	}
	return account.Balance, nil
}

// Deposit adds money to the account
func (r *AccountRepository) Deposit(amount float64) error {
	var account models.Account
	err := r.DB.First(&account).Error
	if err != nil {
		return err
	}

	// Check deposit limits
	today := time.Now().Format("2006-01-02")
	dailyTotal, count, err := r.GetDailyDepositSummary(today)
	if err != nil {
		return err
	}

	if amount > 40000 {
		return errors.New("exceeded maximum deposit per transaction")
	}
	if dailyTotal+amount > 150000 {
		return errors.New("exceeded maximum deposit per day")
	}
	if count >= 4 {
		return errors.New("exceeded maximum deposit frequency per day")
	}

	// Update balance
	err = r.DB.Transaction(func(tx *gorm.DB) error {
		account.Balance += amount
		if err := tx.Save(&account).Error; err != nil {
			return err
		}

		transaction := models.Transaction{
			AccountID: account.ID,
			Type:      "deposit",
			Amount:    amount,
		}
		return tx.Create(&transaction).Error
	})

	return err
}

// Withdraw deducts money from the account
func (r *AccountRepository) Withdraw(amount float64) error {
	var account models.Account
	err := r.DB.First(&account).Error
	if err != nil {
		return err
	}

	// Check withdrawal limits
	today := time.Now().Format("2006-01-02")
	dailyTotal, count, err := r.GetDailyWithdrawalSummary(today)
	if err != nil {
		return err
	}

	if amount > 20000 {
		return errors.New("exceeded maximum withdrawal per transaction")
	}
	if dailyTotal+amount > 50000 {
		return errors.New("exceeded maximum withdrawal per day")
	}
	if count >= 3 {
		return errors.New("exceeded maximum withdrawal frequency per day")
	}
	if account.Balance < amount {
		return errors.New("insufficient funds")
	}

	// Update balance
	err = r.DB.Transaction(func(tx *gorm.DB) error {
		account.Balance -= amount
		if err := tx.Save(&account).Error; err != nil {
			return err
		}

		transaction := models.Transaction{
			AccountID: account.ID,
			Type:      "withdraw",
			Amount:    amount,
		}
		return tx.Create(&transaction).Error
	})

	return err
}

// GetDailyDepositSummary returns total deposits and count for today
func (r *AccountRepository) GetDailyDepositSummary(date string) (float64, int, error) {
	var total float64
	var count int64

	var result struct {
		Total float64
		Count int64
	}

	err := r.DB.Model(&models.Transaction{}).
		Where("type = ? AND DATE(created_at) = ?", "deposit", date).
		Select("SUM(amount) as total, COUNT(*) as count").
		Scan(&result).Error

	total = result.Total
	count = result.Count

	return total, int(count), err
}

// GetDailyWithdrawalSummary returns total withdrawals and count for today
func (r *AccountRepository) GetDailyWithdrawalSummary(date string) (float64, int, error) {
	var total float64
	var count int64

	var result struct {
		Total float64
		Count int64
	}

	err := r.DB.Model(&models.Transaction{}).
		Where("type = ? AND DATE(created_at) = ?", "withdraw", date).
		Select("SUM(amount) as total, COUNT(*) as count").
		Scan(&result).Error

	total = result.Total
	count = result.Count

	return total, int(count), err
}
