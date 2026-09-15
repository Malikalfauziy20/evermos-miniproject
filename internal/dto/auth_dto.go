package dto

type RegisterRequest struct {
	Nama       string `json:"nama" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	NoTelepon  string `json:"no_telepon" validate:"required"`
	Password   string `json:"kata_sandi" validate:"required,min=6"`
	Pekerjaan  string `json:"pekerjaan"`
	IDProvinsi string `json:"id_provinsi"`
	IDKota     string `json:"id_kota"`
}

type LoginRequest struct {
	NoTelepon string `json:"no_telepon" validate:"required"`
	Password  string `json:"kata_sandi" validate:"required"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}
