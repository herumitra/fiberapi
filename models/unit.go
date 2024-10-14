package models

type Unit struct {
	ID   string `gorm:"type:varchar(19);primaryKey" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
}
