package models

import (
	"time"
)

type Subscription struct {
	ID       					string   		`form:"id" json:"id"`
	Email     					string			`form:"email" json:"email"`
	Paket 						int				`form:"paket" json:"paket"`
	NoCC						int				`form:"no_cc" json:"genre"`
	MasaBerlaku					string			`form:"masa_berlaku" json:"masa_berlaku"`
	KodeCVC						int				`form:"kode_cvc" json:"kode_cvc"`
	TanggalLangganan			time.Time		`form:"tgl_langganan" json:"tgl_langganan"`
	TanggalBerhenti				string			`form:"tgl_berhenti" json:"tgl_berhenti"`
}

type SubscriptionResponse struct {
	Status  int  			 `form:"status" json:"status"`
	Message string			 `form:"message" json:"message"`
	Data    []Subscription	 `form:"data" json:"data"`
}