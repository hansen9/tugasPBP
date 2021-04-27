package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"
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

//Signin
func Signin(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	email := r.URL.Query()["email"]
	password := r.URL.Query()["password"]

	row := db.QueryRow("SELECT * FROM user WHERE email=? AND password=?", email[0], password[0])

	var user models.User
	if err := row.Scan(&user.Email, &user.Name, &user.Password, &user.TanggalLahir,
		&user.JenisKelamin, &user.AsalNegara, &user.Status, &user.TipeUser); err != nil {
		sendErrorResponse(w)
	} else {
		generateToken(w, user.Email, user.Name, user.TipeUser)
		sendSuccessResponse(w)
	}
}

// Signout
func Signout(w http.ResponseWriter, r *http.Request) {
	resetUserToken(w)

	sendSuccessResponse(w)
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

// Update member profile
func UpdateMember(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	email := getID(r)
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

// Subscribing
func Subscribe(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	email := getID(r)

	paket := r.Form.Get("paket")
	no_cc, _ := strconv.Atoi(r.Form.Get("no_cc"))
	masa_berlaku := r.Form.Get("masa_berlaku")
	kode_cvc, _ := strconv.Atoi(r.Form.Get("kode_cvc"))
	tgl_langganan := time.Now()
	tgl_berhenti := ""

	_, errQuery := db.Exec("INSERT INTO subscription(email_member,paket,no_cc,masa_berlaku,kode_cvc,tgl_langganan,tgl_berhenti) VALUES (?,?,?,?,?,?,?)",
		email,
		paket,
		no_cc,
		masa_berlaku,
		kode_cvc,
		tgl_langganan,
		tgl_berhenti,
	)

	if errQuery == nil {
		sendSubscriptionSuccessResponse(w, nil)
	} else {
		sendErrorResponse(w)
	}
}

// Stop subscribing
func StopSubscribe(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	email := getID(r)
	tgl_berhenti := time.Now()

	_, errQuery := db.Exec("UPDATE subscription SET tgl_berhenti=? WHERE email_member=?",
		tgl_berhenti,
		email,
	)

	if errQuery == nil {
		sendSubscriptionSuccessResponse(w, nil)
	} else {
		sendErrorResponse(w)
	}
}

// Watch
func Watch(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	id_film := r.Form.Get("id_film")
	email := getID(r)
	tgl_menonton := time.Now()

	_, errQuery := db.Exec("INSERT INTO history(email_member,id_film,tgl_menonton) VALUES (?,?,?)",
		email,
		id_film,
		tgl_menonton,
	)

	if errQuery == nil {
		sendHistorySuccessResponse(w, nil)
	} else {
		sendErrorResponse(w)
	}
}

// Show history
func GetHistory(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	email := getID(r)

	query := "SELECT history.email_member, history.id_film, history.tgl_menonton, " +
		"film.judul, film.tahun, film.genre, film.sutradara, film.pemain_utama, film.sinopsis " +
		"FROM history " +
		"JOIN film ON history.id_film = film.id_film WHERE history.email_member='" + email + "'"
	
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var history models.History
	var histories []models.History
	for rows.Next() {
		if err := rows.Scan(&history.Email, &history.Film.ID, &history.TanggalMenonton,
			&history.Film.Judul, &history.Film.Tahun, &history.Film.Genre,
			&history.Film.Sutradara, &history.Film.PemainUtama, &history.Film.Sinopsis); err != nil {
			sendErrorResponse(w)
			log.Fatal(err.Error())
		} else {
			histories = append(histories, history)
		}
	}

	if len(histories) > 0 {
		sendHistorySuccessResponse(w, histories)
	} else {
		sendErrorResponse(w)
	}
}
