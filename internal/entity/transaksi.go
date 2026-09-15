package entity

import "time"

type Transaksi struct {
	ID           uint `gorm:"primaryKey"`
	IDUser       uint `gorm:"not null"`
	AlamatKirim  uint `gorm:"not null"`
	HargaTotal   int
	KodeInvoice  string
	MethodBayar  string
	CreatedAt    time.Time

	DetailTransaksi []DetailTransaksi `gorm:"foreignKey:IDTransaksi"`
}