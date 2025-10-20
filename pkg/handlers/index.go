package handlers

import (
	"net/http"
	"text/template"

	"github.com/lukegrn/days/pkg/db"
)

func Index(w http.ResponseWriter, r *http.Request) {
	images, err := db.Get().GetAllImages()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Whoops, something went wrong!"))
		return
	}

	t, err := template.ParseFiles("templates/base.html", "templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Whoops, something went wrong!"))
		return
	}

	t.Execute(w, images)
}
