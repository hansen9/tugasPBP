package controllers

import (
	// "encoding/json"
	// "log"
	"net/http"

	"github.com/gorilla/mux"
)

// Register...
func Register(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	email := r.Form.Get("email")
	nama := r.Form.Get("nama")
	password := r.Form.Get("password")
	tgl_lahir := r.Form.Get("tgl_lahir")
	jns_kelamin := r.Form.Get("jns_kelamin")
	asal_negara := r.Form.Get("asal_negara")
	status := "Aktif"
	tipe_user := 0

	_, errQuery := db.Exec("INSERT INTO user(email,nama,password,tgl_lahir,jns_kelamin,asal_negara,status,tipe_user) VALUES (?,?,?,?,?,?,?,?)",
		email,
		nama,
		password,
		tgl_lahir,
		jns_kelamin,
		asal_negara,
		status,
		tipe_user,
	)

	if errQuery == nil {
		sendFilmSuccessResponse(w, nil)
	} else {
		sendErrorResponse(w)
	}
}

// Update
func UpdateMember(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	email := vars["email"]
	nama := r.Form.Get("nama")
	tgl_lahir := r.Form.Get("tgl_lahir")
	jns_kelamin := r.Form.Get("jns_kelamin")

	_, errQuery := db.Exec("UPDATE user SET nama=?, tgl_lahir=?, jns_kelamin=? WHERE tipe_user=0 AND email=?",
		nama,
		tgl_lahir,
		jns_kelamin,
		email,
	)

	if errQuery == nil {
		sendUserSuccessResponse(w, nil)
	} else {
		sendErrorResponse(w)
	}
}
