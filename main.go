package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {
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
	fmt.Println("db connected")

	http.HandleFunc("/", MainInterface)

	http.HandleFunc("/register", RegInterface)
	http.HandleFunc("/register/post", RegUser)

	http.HandleFunc("/login", LogInterface)
	http.HandleFunc("/login/post", LogUser)

	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalln(err)
	}
}