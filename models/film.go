package models

type Film struct {
	ID       		string   	`form:"id" json:"id"`
	Judul     		string		`form:"judul" json:"judul"`
	Tahun 			int			`form:"tahun" json:"tahun"`
	Genre			string		`form:"genre" json:"genre"`
	Sutradara		string		`form:"sutradara" json:"sutradara"`
	PemainUtama		string		`form:"pemain_utama" json:"pemain_utama"`
	Sinopsis		string		`form:"sinopsis" json:"sinopsis"`
}

type FilmResponse struct {
	Status  int    `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
	Data    []Film `form:"data" json:"data"`
}
