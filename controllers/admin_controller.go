package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	models "github.com/tubes/models"
)

// GetMemberByEmail...
func GetMemberByEmail(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM user"

	email := r.URL.Query()["email"]
	if email != nil {
		query += " WHERE email='" + email[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var user models.User
	var users []models.User
	for rows.Next() {
		if err := rows.Scan(&user.Email, &user.Name, &user.Password, &user.TanggalLahir,
			&user.JenisKelamin, &user.AsalNegara, &user.Status, &user.TipeUser); err != nil {
			sendErrorResponse(w)
		} else {
			users = append(users, user)
			sendUserSuccessResponse(w, users)
		}
	}
}

// UpdateUser...
func SuspendMember(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	email := vars["email"]

	_, errQuery := db.Exec("UPDATE user SET status='Terkunci' WHERE tipe_user=0 AND email=?",
		email,
	)

	if errQuery == nil {
		sendUserSuccessResponse(w, nil)
	} else {
		sendErrorResponse(w)
	}
}

//AdminGetFilm
func AdminGetFilm(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM film"

	id := r.URL.Query()["id"]
	judul := r.URL.Query()["judul"]

	if id != nil {
		query += " WHERE id_film = " + id[0]
	} else if judul != nil {
		query += " WHERE judul LIKE '%" + judul[0] + "%'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var film models.Film
	var films []models.Film
	for rows.Next() {
		if err := rows.Scan(&film.ID, &film.Judul, &film.Tahun, &film.Genre, &film.Sutradara, &film.PemainUtama, &film.Sinopsis); err != nil {
			sendErrorResponse(w)
		} else {
			films = append(films, film)

		}
	}

	if len(films) > 0 {
		sendFilmSuccessResponse(w, films)
	} else {
		sendErrorResponse(w)
	}
}

// InsertFilm...
func InsertFilm(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	judul := r.Form.Get("judul")
	tahun, _ := strconv.Atoi(r.Form.Get("tahun"))
	genre := r.Form.Get("genre")
	sutradara := r.Form.Get("sutradara")
	pemain_utama := r.Form.Get("pemain_utama")
	sinopsis := r.Form.Get("sinopsis")

	_, errQuery := db.Exec("INSERT INTO film(judul, tahun, genre, sutradara, pemain_utama, sinopsis) VALUES (?,?,?,?,?,?)",
		judul,
		tahun,
		genre,
		sutradara,
		pemain_utama,
		sinopsis,
	)

	if errQuery == nil {
		sendFilmSuccessResponse(w, nil)
	} else {
		sendErrorResponse(w)
	}
}

// UpdateFilmById...
func UpdateFilmById(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	judul := r.Form.Get("judul")
	tahun, _ := strconv.Atoi(r.Form.Get("tahun"))
	genre := r.Form.Get("genre")
	sutradara := r.Form.Get("sutradara")
	pemain_utama := r.Form.Get("pemain_utama")
	sinopsis := r.Form.Get("sinopsis")

	_, errQuery := db.Exec("UPDATE film SET judul=?, tahun=?, genre=?, sutradara=?, pemain_utama=?, sinopsis=? WHERE id_film=?",
		judul,
		tahun,
		genre,
		sutradara,
		pemain_utama,
		sinopsis,
		id,
	)

	if errQuery == nil {
		sendFilmSuccessResponse(w, nil)
	} else {
		sendErrorResponse(w)
	}
}
