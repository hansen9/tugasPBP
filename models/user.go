package models

type User struct {
	Email       	string   	 `form:"id" json:"id"`
	Name     		string		`form:"name" json:"name"`
	Password 		string		`form:"password" json:"password"`
	TanggalLahir	string		`form:"tanggal_lahir" json:"tanggal_lahir"`
	JenisKelamin	string		`form:"jenis_kelamin" json:"jenis_kelamin"`
	AsalNegara		string		`form:"asal_negara" json:"asal_negara"`
	Status			string		`form:"status" json:"status"`
	TipeUser		int			`form:"tipe_user" json:"tipe_user"`
}

type UserResponse struct {
	Status  int    `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}

type ErrorResponse struct {
	Status  int    `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
}