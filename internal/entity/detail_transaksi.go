package entity

type DetailTransaksi struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	IDTransaksi uint `gorm:"not null" json:"id_transaksi"`
	IDLogProduk uint `gorm:"not null" json:"id_log_produk"`
	IDToko      uint `gorm:"not null" json:"id_toko"`
	Kuantitas   int  `json:"kuantitas"`

	LogProduk LogProduk `gorm:"foreignKey:IDLogProduk" json:"log_produk,omitempty"`
}

func (DetailTransaksi) TableName() string {
	return "detail_transaksis"
}
