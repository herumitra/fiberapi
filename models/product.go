package models

type Product struct {
	ID               string `gorm:"type:varchar(19);primaryKey" json:"id"`
	Name             string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	Price            int    `gorm:"type:int;not null" json:"price" validate:"required"`
	Description      string `gorm:"type:text;not null" json:"description" validate:"required"`
	PlaceId          string `gorm:"type:varchar(19);not null" json:"place_id" validate:"required"`
	UnitId           string `gorm:"type:varchar(19);not null" json:"unit_id" validate:"required"`
	CategoriId       string `gorm:"type:varchar(19);not null" json:"categori_id" validate:"required"`
	SKUCode          string `gorm:"type:varchar(100);not null" json:"sku_code" validate:"required"`
	BarCode          string `gorm:"type:varchar(100);not null" json:"bar_code" validate:"required"`
	PurchasePrice    int    `gorm:"type:int(25);not null;default:0" json:"purchase_price" validate:"required"`
	SellingPrice     int    `gorm:"type:int(25);not null;default:0" json:"selling_price" validate:"required"`
	RecomendPrice    int    `gorm:"type:int(25);not null;default:0" json:"recomend_price" validate:"required"`
	GroceryPrice     int    `gorm:"type:int(25);default:0" json:"grocery_price"`
	AlternatePrice   int    `gorm:"type:int(25);default:0" json:"alternate_price"`
	TaxAble          string `gorm:"type:enum('yes','no');default:'yes'" json:"tax_able" validate:"required"`
	ActiveIngredient string `gorm:"type:varchar(255);default:null" json:"active_ingredient"`
	BranchId         string `gorm:"type:varchar(19);not null" json:"branch_id" validate:"required"`
	ItemStatus       string `gorm:"type:enum('active','inactive');default:'active'" json:"item_status" validate:"required"`
	// Category    string `gorm:"type:enum('OBAT BATUK','OBAT PILEK','OBAT DEMAM','OBAT PANAS','OBAT PENCERNAAN','OBAT SESAK NAPAS','OBAT MASUK ANGIN','OBAT MATA','OBAT KUMUR','OBAT LUAR','OBAT KULIT','OBAT TETES','SIRUP','DROP','SUPLEMEN','HERBAL','MAKANAN','MINUMAN','SUSU BUBUK','JASA','ALKES','PEMBALUT','POPOK BAYI','POPOK DEWASA','UNDERPAD','REPRODUKSI','LAINNYA');not null" json:"category" validate:"required"`
}
