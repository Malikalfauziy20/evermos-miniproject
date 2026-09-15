package entity

type Produk struct {
	ID           uint    `gorm:"primaryKey"`
	NamaProduk   string  `gorm:"size:255;not null"`
	Slug         string
	HargaReseller string
	HargaKonsumen string
	Stok         int
	Deskripsi    string
	IDToko       uint `gorm:"not null"`
	IDCategory   uint `gorm:"not null"`

	Toko     Toko          `gorm:"foreignKey:IDToko"`
	Kategori Kategori      `gorm:"foreignKey:IDCategory"`
	Foto     []FotoProduk  `gorm:"foreignKey:IDProduk"`
}