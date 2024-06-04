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

func RegInterface(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("resources/register.html"))
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
func RegUser(w http.ResponseWriter, r *http.Request) {
	var user database.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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

	var name string
	msg := "User is exist"
	isExist := true
	err = db.QueryRow("SELECT name FROM users WHERE email = $1", user.Email).Scan(&name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			isExist = false
		} else {
			log.Fatalln(err)
		}
	}

	if !isExist {
		msg = "User adding is success"
		hashpass, err := user.HashPassword()
		if err != nil {
			log.Fatalln(err)
		}
		_, err = db.Exec("INSERT INTO users (email, name, hashpass) VALUES ($1, $2, $3)", user.Email, user.Name, hashpass)
		if err != nil {
			log.Fatalln(err)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{"message": msg})
	if err != nil {
		return
	}
}
