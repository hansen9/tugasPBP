package models

type History struct {
	Email           string `form:"email" json:"email"`
	Film            Film   `form:"film" json:"film"`
	TanggalMenonton string `form:"tgl_menonton" json:"tgl_menonton"`
}

type HistoryResponse struct {
	Status  int       `form:"status" json:"status"`
	Message string    `form:"message" json:"message"`
	Data    []History `form:"data" json:"data"`
}
