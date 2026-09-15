package entity

type FotoProduk struct {
	ID       uint `gorm:"primaryKey"`
	IDProduk uint `gorm:"not null"`
	URL      string
}