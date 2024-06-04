package handlers

import (
	"AuthGo/database"
	"database/sql"
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
)

func LogInterface(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("resources/login.html"))
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
func LogUser(w http.ResponseWriter, r *http.Request) {
	var user database.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalln(err)
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalln(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)

	var hashpass string
	msg := "Password incorrect"
	err = db.QueryRow("SELECT hashpass FROM users WHERE email = $1", user.Email).Scan(&hashpass)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			msg = "No user with this email"
		} else {
			log.Fatalln(err)
		}
	}
	cmpPass := database.ComparePassword(hashpass, user.Pass)
	if cmpPass {
		msg = "Logon success"
		// Realize logic with cookies
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{"message": msg})
	if err != nil {
		return
	}
}
