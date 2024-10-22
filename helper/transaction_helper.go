package helper

import (
	"math/rand"
	"simple_bank_solid/model/domain"

	"gorm.io/gorm"
)

const charset = "0123456789"

func Generate(identifier string) string {
	var length int = 10

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return identifier + string(b)
}

func ValidAccount(tx *gorm.DB, accountID int64, currency string) (domain.Account, bool) {
	account := domain.Account{}
	err := tx.Model(&domain.Account{}).Take(&account, "id = ?", accountID).Error
	if err != nil {
		return account, false
	}
	if account.Currency != currency {
		return account, false
	}
	return account, true
}
