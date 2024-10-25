package models

// MemberCategory model
type MemberCategory struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	BranchId string `gorm:"type:varchar(19);not null" json:"branch_id" validate:"required"`
}

// Member model
type Member struct {
	ID               string `gorm:"type:varchar(19);primaryKey" json:"id" validate:"required"`
	Name             string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	Phone            string `gorm:"type:varchar(100);" json:"phone"`
	Address          string `gorm:"type:text;" json:"address"`
	MemberCategoryId uint   `gorm:"not null" json:"member_category_id" validate:"required"`
	Saldo            int    `gorm:"type:int(25);not null;default:0" json:"saldo" validate:"required"`
	BranchId         string `gorm:"type:varchar(19);not null" json:"branch_id" validate:"required"`
}

// Unit model
type Unit struct {
	ID       string `gorm:"type:varchar(19);primaryKey" json:"id" validate:"required"`
	Name     string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	BranchId string `gorm:"type:varchar(19);not null" json:"branch_id" validate:"required"`
}

// UnitConversion model
type UC struct {
	ID          string `gorm:"type:varchar(19);primaryKey" json:"id" validate:"required"`
	UnitInitId  string `gorm:"type:varchar(19);not null" json:"unit_init_id" validate:"required"`
	UnitFinalId string `gorm:"type:varchar(19);not null" json:"unit_final_id" validate:"required"`
	ValueConv   int    `gorm:"type:int(25);not null;default:0" json:"value_conv" validate:"required"`
	BranchId    string `gorm:"type:varchar(19);not null" json:"branch_id" validate:"required"`
}

// Place model
type Place struct {
	ID       string `gorm:"type:varchar(19);primaryKey" json:"id" validate:"required"`
	Name     string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	BranchId string `gorm:"type:varchar(19);not null" json:"branch_id" validate:"required"`
}

// Item model
type Item struct {
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
	QtyStock         int    `gorm:"type:int(25);not null;default:0" json:"qty_stock" validate:"required"`
	ItemStatus       string `gorm:"type:enum('active','inactive');default:'active'" json:"item_status" validate:"required"`
	BranchId         string `gorm:"type:varchar(19);not null" json:"branch_id" validate:"required"`
	// Category    string `gorm:"type:enum('OBAT BATUK','OBAT PILEK','OBAT DEMAM','OBAT PANAS','OBAT PENCERNAAN','OBAT SESAK NAPAS','OBAT MASUK ANGIN','OBAT MATA','OBAT KUMUR','OBAT LUAR','OBAT KULIT','OBAT TETES','SIRUP','DROP','SUPLEMEN','HERBAL','MAKANAN','MINUMAN','SUSU BUBUK','JASA','ALKES','PEMBALUT','POPOK BAYI','POPOK DEWASA','UNDERPAD','REPRODUKSI','LAINNYA');not null" json:"category" validate:"required"`
}

// ItemCategory model
type ItemCategory struct {
	ID       string `gorm:"type:varchar(19);primaryKey" json:"id" validate:"required"`
	Name     string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	BranchId string `gorm:"type:varchar(19);not null" json:"branch_id" validate:"required"`
}

// Supplier model
type Supplier struct {
	ID       string `gorm:"type:varchar(19);primaryKey" json:"id" validate:"required"`
	Name     string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	Address  string `gorm:"type:text;" json:"address"`
	Phone    string `gorm:"type:varchar(100);" json:"phone"`
	Email    string `gorm:"type:varchar(100);" json:"email"`
	BranchId string `gorm:"type:varchar(19);not null" json:"branch_id" validate:"required"`
}
