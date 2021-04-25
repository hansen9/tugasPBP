package controllers

import (
	// "encoding/json"
	// "log"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	models "github.com/tubes/models"
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

//GetFilm
func GetFilm(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM film"

	judul := r.URL.Query()["judul"]
	tahun := r.URL.Query()["tahun"]
	genre := r.URL.Query()["genre"]
	pemain_utama := r.URL.Query()["pemain_utama"]
	sutradara := r.URL.Query()["sutradara"]
	sinopsis := r.URL.Query()["sinopsis"]
	if judul != nil {
		query += " WHERE judul = '" + judul[0] + "'"
		if tahun != nil {
			query += " AND tahun = '" + tahun[0] + "'"
		}
		if genre != nil {
			query += " AND genre = '" + genre[0] + "'"
		}
		if pemain_utama != nil {
			query += " AND pemain_utama = '" + pemain_utama[0] + "'"
		}
		if sutradara != nil {
			query += " AND sutradara = '" + sutradara[0] + "'"
		}
		if sinopsis != nil {
			query += " AND sinopsis = '" + sinopsis[0] + "'"
		}
	} else if tahun != nil {
		query += " WHERE tahun = '" + tahun[0] + "'"
		if genre != nil {
			query += " AND genre = '" + genre[0] + "'"
		}
		if pemain_utama != nil {
			query += " AND pemain_utama = '" + pemain_utama[0] + "'"
		}
		if sutradara != nil {
			query += " AND sutradara = '" + sutradara[0] + "'"
		}
		if sinopsis != nil {
			query += " AND sinopsis = '" + sinopsis[0] + "'"
		}
	} else if genre != nil {
		query += " WHERE genre = '" + genre[0] + "'"
		if pemain_utama != nil {
			query += " AND pemain_utama = '" + pemain_utama[0] + "'"
		}
		if sutradara != nil {
			query += " AND sutradara = '" + sutradara[0] + "'"
		}
		if sinopsis != nil {
			query += " AND sinopsis = '" + sinopsis[0] + "'"
		}
	} else if pemain_utama != nil {
		query += " WHERE pemain_utama = '" + pemain_utama[0] + "'"
		if sutradara != nil {
			query += " AND sutradara = '" + sutradara[0] + "'"
		}
		if sinopsis != nil {
			query += " AND sinopsis = '" + sinopsis[0] + "'"
		}
	} else if sutradara != nil {
		query += " WHERE sutradara = '" + sutradara[0] + "'"
		if sinopsis != nil {
			query += " AND sinopsis = '" + sinopsis[0] + "'"
		}
	} else if sinopsis != nil {
		query += " WHERE sinopsis = '" + sinopsis[0] + "'"
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
