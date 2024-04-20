package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)

func MainInterface(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("resources/index.html"))
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func RegInterface(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("resources/register.html"))
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
func RegUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := ConnectDB()
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
func LogInterface(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("resources/login.html"))
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
func LogUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalln(err)
	}

	db, err := ConnectDB()
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
	cmpPass := ComparePassword(hashpass, user.Pass)
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
