package app

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("html/09_CSS.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
