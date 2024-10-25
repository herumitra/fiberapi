package models

// Prices model
type Prices struct {
	ID        string `gorm:"type:varchar(19);primaryKey" json:"id" validate:"required"`
	Quantity  int    `gorm:"type:int;not null;default:0" json:"quantity" validate:"required"`
	Price     int    `gorm:"type:int;not null;default:0" json:"price" validate:"required"`
	ItemId    string `gorm:"type:varchar(19);not null" json:"item_id" validate:"required"`
	SpplierId string `gorm:"type:varchar(19);not null" json:"supplier_id" validate:"required"`
	BranchId  string `gorm:"type:varchar(19);not null" json:"branch_id" validate:"required"`
}
