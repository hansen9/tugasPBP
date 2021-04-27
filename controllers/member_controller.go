package controllers

import (
	// "encoding/json"
	// "log"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	models "github.com/tubes/models"
)

// Register (insert member)
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

// Search film for member
func MemberGetFilm(w http.ResponseWriter, r *http.Request) {
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
		query += " WHERE judul LIKE '%" + judul[0] + "%'"
		if tahun != nil {
			query += " AND tahun LIKE '%" + tahun[0] + "%'"
		}
		if genre != nil {
			query += " AND genre LIKE '%" + genre[0] + "%'"
		}
		if pemain_utama != nil {
			query += " AND pemain_utama LIKE '%" + pemain_utama[0] + "%'"
		}
		if sutradara != nil {
			query += " AND sutradara LIKE '%" + sutradara[0] + "%'"
		}
		if sinopsis != nil {
			query += " AND sinopsis LIKE '%" + sinopsis[0] + "%'"
		}
	} else if tahun != nil {
		query += " WHERE tahun LIKE '%" + tahun[0] + "%'"
		if genre != nil {
			query += " AND genre LIKE '%" + genre[0] + "%'"
		}
		if pemain_utama != nil {
			query += " AND pemain_utama LIKE '%" + pemain_utama[0] + "%'"
		}
		if sutradara != nil {
			query += " AND sutradara LIKE '%" + sutradara[0] + "%'"
		}
		if sinopsis != nil {
			query += " AND sinopsis LIKE '%" + sinopsis[0] + "%'"
		}
	} else if genre != nil {
		query += " WHERE genre LIKE '%" + genre[0] + "%'"
		if pemain_utama != nil {
			query += " AND pemain_utama LIKE '%" + pemain_utama[0] + "%'"
		}
		if sutradara != nil {
			query += " AND sutradara LIKE '%" + sutradara[0] + "%'"
		}
		if sinopsis != nil {
			query += " AND sinopsis LIKE '%" + sinopsis[0] + "%'"
		}
	} else if pemain_utama != nil {
		query += " WHERE pemain_utama LIKE '%" + pemain_utama[0] + "%'"
		if sutradara != nil {
			query += " AND sutradara LIKE '%" + sutradara[0] + "%'"
		}
		if sinopsis != nil {
			query += " AND sinopsis LIKE '%" + sinopsis[0] + "%'"
		}
	} else if sutradara != nil {
		query += " WHERE sutradara LIKE '%" + sutradara[0] + "%'"
		if sinopsis != nil {
			query += " AND sinopsis LIKE '%" + sinopsis[0] + "%'"
		}
	} else if sinopsis != nil {
		query += " WHERE sinopsis LIKE '%" + sinopsis[0] + "%'"
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

// Update member
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

// Watch
func Watch(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	err := r.ParseForm()

	id_film := r.Form.Get("id_film")
	emailUser := r.Form.Get("email")
	tgl_menonton := r.Form.Get("tanggal_menonton")

	_, errQuery := db.Exec("INSERT INTO history(email_member,id_film,tgl_menonton) VALUES (?,?,?)",
		emailUser,
		id_film,
		tgl_menonton,
	)

	query := "SELECT * FROM film WHERE id_film = ?"
	rows, err := db.Query(query, id_film)

	var film models.Film
	var films []models.Film

	for rows.Next() {
		if err := rows.Scan(&film.ID, &film.Judul, &film.Tahun, &film.Genre, &film.Sutradara, &film.PemainUtama, &film.Sinopsis); err != nil {
			log.Fatal(err.Error)
		} else {
			films = append(films, film)
		}
	}

	if err == nil && errQuery == nil {
		sendFilmSuccessResponse(w, films)
	} else {
		_, errQuery := db.Exec("UPDATE history SET tgl_menonton = ? WHERE email_member = ? AND id_film = ?",
			tgl_menonton,
			emailUser,
			id_film,
		)
		if errQuery == nil {
			sendFilmSuccessResponse(w, films)
		} else {
			sendErrorResponse(w)
		}
	}
}

//riwayat
func ShowHistory(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT u.email,f.id_film,f.judul,f.tahun,f.genre,f.sutradara,f.pemain_utama,f.sinopsis,h.tgl_menonton FROM history h INNER JOIN film f ON f.id_film = h.id_film INNER JOIN user u ON u.email = h.email_member"

	email := r.URL.Query()["email"]
	if email != nil {
		query += " AND u.email='" + email[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var history models.History
	var histories []models.History

	for rows.Next() {
		if err := rows.Scan(&history.Email, &history.Film.ID, &history.Film.Judul, &history.Film.Tahun,
			&history.Film.Genre, &history.Film.Sutradara, &history.Film.PemainUtama, &history.Film.Sinopsis,
			&history.TanggalMenonton); err != nil {
			sendErrorResponse(w)
		} else {
			histories = append(histories, history)
			sendHistorySuccessResponse(w, histories)
		}
	}
}
