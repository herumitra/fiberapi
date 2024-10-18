package models

type Product struct {
	ID          string `gorm:"type:varchar(19);primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	Price       int    `gorm:"type:int;not null" json:"price" validate:"required"`
	Description string `gorm:"type:text;not null" json:"description" validate:"required"`
	PlaceId     string `gorm:"type:varchar(19);not null" json:"place_id" validate:"required"`
	UnitId      string `gorm:"type:varchar(19);not null" json:"unit_id" validate:"required"`
	Category    string `gorm:"type:enum('OBAT BATUK','OBAT PILEK','OBAT DEMAM','OBAT PANAS','OBAT PENCERNAAN','OBAT SESAK NAPAS','OBAT MASUK ANGIN','OBAT MATA','OBAT KUMUR','OBAT LUAR','OBAT KULIT','OBAT TETES','SIRUP','DROP','SUPLEMEN','HERBAL','MAKANAN','MINUMAN','SUSU BUBUK','JASA','ALKES','PEMBALUT','POPOK BAYI','POPOK DEWASA','UNDERPAD','REPRODUKSI','LAINNYA');not null" json:"category" validate:"required"`
	BranchId    string `gorm:"type:varchar(19);not null" json:"branch_id" validate:"required"`
}
