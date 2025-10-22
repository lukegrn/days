package handlers

import (
	"net/http"
	"text/template"
)

func Robots(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/robots.txt")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Whoops, something went wrong!"))
		return
	}

	t.Execute(w, nil)
}
