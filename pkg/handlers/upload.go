package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"

	"github.com/lukegrn/days/pkg/db"
	"github.com/lukegrn/days/pkg/img"
)

func ShowUpload(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/base.html", "templates/upload.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Whoops, something went wrong!"))
		return
	}

	t.Execute(w, nil)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	f, _, err := r.FormFile("picture")
	defer f.Close()

	caption := r.Form.Get("caption")
	date := r.Form.Get("date")

	filename := date + ".jpg"
	preview_filename := date + "-resized.jpg"

	if caption == "" || date == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing request param"))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Whoops, bad request"))
		return
	}

	photoBuf := bytes.NewBuffer(nil)
	if _, err := io.Copy(photoBuf, f); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Whoops, something went wrong!"))
		return
	}

	// Save jpg
	err = os.WriteFile("./static/"+filename, photoBuf.Bytes(), 0644)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Whoops, failed to save image!"))
		return
	}

	// Save resized preview
	preview, err := os.Create("./static/" + preview_filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Whoops, failed to open file for resized output!"))
		return
	}

	err = img.CreatePreview(photoBuf, preview)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Whoops, failed to create resized preview file!"))
		return
	}

	err = db.Get().PutImage(date, filename, caption)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Whoops, failed to save image info to database: %s", err.Error())))
		return
	}

	// Success
	w.Write([]byte(fmt.Sprintf("caption: %s\n", caption)))
	w.Write([]byte(fmt.Sprintf("date: %s\n", date)))
	w.Write([]byte("Success!"))

}
