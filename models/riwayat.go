package models

type Riwayat struct {
	Email       			string   	`form:"email" json:"email"`
	Film     				Film		`form:"film" json:"film"`
	TanggalMenonton 		string		`form:"tgl_menonton" json:"tgl_menonton"`
}

type RiwayatResponse struct {
	Status  int    		`form:"status" json:"status"`
	Message string 		`form:"message" json:"message"`
	Data    []Riwayat	`form:"data" json:"data"`
}
