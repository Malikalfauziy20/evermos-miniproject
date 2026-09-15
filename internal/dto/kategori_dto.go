package dto

type CreateKategoriRequest struct {
	NamaKategori string `json:"nama_kategori" validate:"required"`
}

type UpdateKategoriRequest struct {
	NamaKategori string `json:"nama_kategori" validate:"required"`
}
