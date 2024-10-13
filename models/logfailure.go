package models

type LogFailure struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string `gorm:"type:varchar(100);not null" json:"username"`
	Timestamp string `gorm:"type:varchar(100);not null" json:"timestamp"`
}
