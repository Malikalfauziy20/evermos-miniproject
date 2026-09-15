package entity

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Nama         string `gorm:"size:255;not null"`
	Email        string `gorm:"size:255;unique;not null"`
	NoTelepon    string `gorm:"size:20;unique;not null"`
	Password     string `gorm:"not null" json:"-"`
	TanggalLahir time.Time
	Pekerjaan    string
	IsAdmin      bool `gorm:"default:false"`
	IDProvinsi   string
	IDKota       string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Toko   Toko     `gorm:"foreignKey:IDUser"`
	Alamat []Alamat `gorm:"foreignKey:IDUser"`
}
