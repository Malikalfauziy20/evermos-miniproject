package dto

type CreateProdukRequest struct {
	NamaProduk    string `json:"nama_produk" validate:"required"`
	HargaReseller string `json:"harga_reseller" validate:"required"`
	HargaKonsumen string `json:"harga_konsumen" validate:"required"`
	Stok          int    `json:"stok" validate:"required"`
	Deskripsi     string `json:"deskripsi"`
	IDCategory    uint   `json:"id_category" validate:"required"`
}

type UpdateProdukRequest struct {
	NamaProduk    string `json:"nama_produk"`
	HargaReseller string `json:"harga_reseller"`
	HargaKonsumen string `json:"harga_konsumen"`
	Stok          *int   `json:"stok"`
	Deskripsi     string `json:"deskripsi"`
	IDCategory    uint   `json:"id_category"`
}

type ProdukFilter struct {
	Page       int
	Limit      int
	NamaProduk string
	CategoryID uint
}
