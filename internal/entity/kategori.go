package entity

type Kategori struct {
	ID           uint   `gorm:"primaryKey"`
	NamaKategori string `gorm:"size:255;not null"`
}