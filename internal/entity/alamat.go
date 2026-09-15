package entity

type Alamat struct {
	ID           uint `gorm:"primaryKey"`
	IDUser       uint `gorm:"not null"`
	JudulAlamat  string
	NamaPenerima string
	NoTelp       string
	DetailAlamat string
}