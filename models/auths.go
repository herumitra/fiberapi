package models

import (
	"time"

	"gorm.io/gorm"
)

// User model
type User struct {
	ID          string `gorm:"type:varchar(19);primaryKey" json:"id" validate:"required"`
	Username    string `gorm:"type:varchar(100);unique;not null" json:"username" validate:"required"`
	Password    string `gorm:"type:varchar(100);not null" json:"password" validate:"required"`
	BranchId    string `gorm:"type:varchar(19);not null" json:"branch_id" validate:"required"`
	Name        string `gorm:"type:varchar(255);not null" json:"name" validate:"required"`
	Address     string `gorm:"type:text;" json:"address"`
	StatusUser  string `gorm:"type:enum('active','inactive');default:'active'" json:"status_user" validate:"required"`
	Authorities string `gorm:"type:enum('user','cashier','finance','administrator');default:'user'" json:"authorities" validate:"required"`
}

// Branch model
type Branch struct {
	ID            string    `gorm:"type:varchar(19);primaryKey" json:"id"`
	Name          string    `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	Address       string    `gorm:"type:varchar(255)" json:"address" validate:"required"`
	Phone         string    `gorm:"type:varchar(100);" json:"phone"`
	Email         string    `gorm:"type:varchar(100);" json:"email"`
	SiaId         string    `gorm:"type:varchar(19);" json:"sia_id"`
	SiaName       string    `gorm:"type:varchar(100);" json:"sia_name"`
	SipaId        string    `gorm:"type:varchar(19);" json:"sipa_id"`
	SipaName      string    `gorm:"type:varchar(100);" json:"sipa_name"`
	PicName       string    `gorm:"type:varchar(100);not null" json:"pic_name" validate:"required"`
	PicId         string    `gorm:"type:varchar(19);not null" json:"pic_id" validate:"required"`
	ApingId       string    `gorm:"type:varchar(19);" json:"aping_id"`
	ApingName     string    `gorm:"type:varchar(100);" json:"aping_name"`
	BankName      string    `gorm:"type:varchar(100);" json:"bank_name"`
	BankId        string    `gorm:"type:varchar(25);" json:"bank_id"`
	BankAccount   string    `gorm:"type:varchar(100);" json:"bank_account"`
	TaxPercentage int       `gorm:"type:int(3);default:0" json:"tax_percentage"`
	JournalMethod string    `gorm:"type:enum('manual','automatic'); default:'automatic'" json:"journal_method" validate:"required"`
	LicenseDate   time.Time `gorm:"not null" json:"expires_at"`
}

// BeforeCreate hook to set LicenseDate with a month increment, handling end of month cases
func (b *Branch) BeforeCreate(tx *gorm.DB) (err error) {
	currentDate := time.Now()
	nextMonth := currentDate.AddDate(0, 1, 0) // Add one month

	// Handle end-of-month overflow by checking the day
	if nextMonth.Day() < currentDate.Day() {
		nextMonth = nextMonth.AddDate(0, 0, -nextMonth.Day()) // Set to last day of next month
	}

	b.LicenseDate = nextMonth
	return nil
}

// LogFailure model
type LogFailure struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar(100);not null" json:"username"`
	Timestamp time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"json:"timestamp"`
}

// TokenBlacklist model
type TokenBlacklist struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Token     string    `gorm:"type:text;not null" json:"token"`
	ExpiresAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"expires_at"`
}
