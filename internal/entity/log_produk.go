package entity

type LogProduk struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	IDProduk      uint   `json:"id_produk"`
	NamaProduk    string `json:"nama_produk"`
	Slug          string `json:"slug"`
	HargaReseller string `json:"harga_reseller"`
	HargaKonsumen string `json:"harga_konsumen"`
	Deskripsi     string `json:"deskripsi"`
	IDToko        uint   `json:"id_toko"`
	IDCategory    uint   `json:"id_category"`
}

func (LogProduk) TableName() string {
	return "log_produks"
}
