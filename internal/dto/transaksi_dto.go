package dto

type DetailItemRequest struct {
	IDProduk  uint `json:"id_produk" validate:"required"`
	Kuantitas int  `json:"kuantitas" validate:"required,min=1"`
}

type CheckoutRequest struct {
	IDAlamat    uint                `json:"id_alamat" validate:"required"`
	DetailTrx   []DetailItemRequest `json:"detail_trx" validate:"required,min=1"`
	MethodBayar string              `json:"method_bayar"`
}
