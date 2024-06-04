package main

import (
	"AuthGo/database"
	"AuthGo/handlers"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {
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
	fmt.Println("db connected")

	http.HandleFunc("/", handlers.MainInterface)

	http.HandleFunc("/register", handlers.RegInterface)
	http.HandleFunc("/register/post", handlers.RegUser)

	http.HandleFunc("/login", handlers.LogInterface)
	http.HandleFunc("/login/post", handlers.LogUser)

	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
