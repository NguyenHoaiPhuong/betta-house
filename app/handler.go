package app

import (
	"net/http"
	"text/template"

	"github.com/NguyenHoaiPhuong/betta-house/utils"
)

func home(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("static/Login.html")
	utils.CheckError(err)

	err = tpl.Execute(w, nil)
	utils.CheckError(err)
}
