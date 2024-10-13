package models

type Branch struct {
	ID      string `gorm:"type:varchar(19);primaryKey" json:"id"`
	Name    string `gorm:"type:varchar(100);not null" json:"name"`
	Address string `gorm:"type:varchar(255)" json:"address"`
}
