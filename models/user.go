package models

type User struct {
	ID          string `gorm:"type:varchar(19);primaryKey" json:"id"`
	Username    string `gorm:"type:varchar(100);unique;not null" json:"username"`
	Password    string `gorm:"type:varchar(100);not null" json:"password"`
	IDBranch    string `gorm:"type:varchar(19);not null" json:"id_branch"`
	StatusUser  string `gorm:"type:enum('active','inactive');default:'active'" json:"status_user"`
	Authorities string `gorm:"type:enum('user','cashier','administrator','finance');default:'user'" json:"authorities"`
}
