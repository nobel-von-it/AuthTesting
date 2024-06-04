package handlers

import (
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
