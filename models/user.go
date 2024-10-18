package models

type User struct {
	ID          string `gorm:"type:varchar(19);primaryKey" json:"id" validate:"required"`
	Username    string `gorm:"type:varchar(100);unique;not null" json:"username" validate:"required"`
	Password    string `gorm:"type:varchar(100);not null" json:"password" validate:"required"`
	BranchId    string `gorm:"type:varchar(19);not null" json:"branch_id" validate:"required"`
	Name        string `gorm:"type:varchar(255);not null" json:"name" validate:"required"`
	Address     string `gorm:"type:text;" json:"address"`
	StatusUser  string `gorm:"type:enum('active','inactive');default:'active'" json:"status_user" validate:"required"`
	Authorities string `gorm:"type:enum('user','cashier','administrator','finance');default:'user'" json:"authorities" validate:"required"`
}
