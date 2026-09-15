package dto

type UpdateUserRequest struct {
	Nama       string `json:"nama"`
	Email      string `json:"email" validate:"omitempty,email"`
	NoTelepon  string `json:"no_telepon"`
	Pekerjaan  string `json:"pekerjaan"`
	IDProvinsi string `json:"id_provinsi"`
	IDKota     string `json:"id_kota"`
}
