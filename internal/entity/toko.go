package entity

type Toko struct {
	ID       uint   `gorm:"primaryKey"`
	IDUser   uint   `gorm:"not null"`
	NamaToko string `gorm:"size:255;not null"`
	URLFoto  string
}