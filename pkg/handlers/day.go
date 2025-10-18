package handlers

import (
	"net/http"
	"text/template"

	"github.com/lukegrn/days/pkg/db"
)

func ShowDay(w http.ResponseWriter, r *http.Request) {
	date := r.PathValue("date")
	img, err := db.Get().GetImage(date)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Whoops, not found!"))
		return
	}

	data := struct {
		Preview string
		Full    string
	}{Preview: img.Path, Full: img.Path}

	t, err := template.ParseFiles("templates/base.html", "templates/img.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Whoops, something went wrong!"))
		return
	}

	t.Execute(w, data)
}
