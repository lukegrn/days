package handlers

import (
	"net/http"
	"text/template"
)

func About(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/base.html", "templates/about.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Whoops, something went wrong!"))
		return
	}

	t.Execute(w, nil)
}
