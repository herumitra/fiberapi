package models

type Branch struct {
	ID            string `gorm:"type:varchar(19);primaryKey" json:"id"`
	Name          string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	Address       string `gorm:"type:varchar(255)" json:"address" validate:"required"`
	Phone         string `gorm:"type:varchar(100);" json:"phone"`
	Email         string `gorm:"type:varchar(100);" json:"email"`
	SiaId         string `gorm:"type:varchar(19);" json:"sia_id"`
	SiaName       string `gorm:"type:varchar(100);" json:"sia_name"`
	SipaId        string `gorm:"type:varchar(19);" json:"sipa_id"`
	SipaName      string `gorm:"type:varchar(100);" json:"sipa_name"`
	PicName       string `gorm:"type:varchar(100);not null" json:"pic_name" validate:"required"`
	PicId         string `gorm:"type:varchar(19);not null" json:"pic_id" validate:"required"`
	ApingId       string `gorm:"type:varchar(19);" json:"aping_id"`
	ApingName     string `gorm:"type:varchar(100);" json:"aping_name"`
	BankName      string `gorm:"type:varchar(100);" json:"bank_name"`
	BankId        string `gorm:"type:varchar(25);" json:"bank_id"`
	BankAccount   string `gorm:"type:varchar(100);" json:"bank_account"`
	JournalMethod string `gorm:"type:enum('manual','automatic'); default:'automatic'" json:"journal_method" validate:"required"`
}
