package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lukegrn/days/pkg/db"
	"github.com/lukegrn/days/pkg/handlers"
	"github.com/lukegrn/days/pkg/middleware"
)

func main() {
	// Conf values
	var port string
	var dbFile string
	var pwHash string

	if port = os.Getenv("DAYS_PORT"); port == "" {
		port = ":8080"
		log.Printf("No port, defaulting to %s", port)
	}

	if dbFile = os.Getenv("DAYS_DB_FILE"); dbFile == "" {
		dbFile = "/tmp/days_db.sqlite"
		log.Printf("No db file path specified, defaulting to %s", dbFile)
	}

	if pwHash = os.Getenv("DAYS_PW"); pwHash == "" {
		log.Fatal("No password hash specified, exiting. Set DAYS_PW and try again.")
	}

	err := db.Get().SetupDB(dbFile)
	if err != nil {
		log.Fatalf("Failed to setup DB: %s", err.Error())
	}
	defer db.Get().Close()

	// Ensure photo dir exists
	err = os.MkdirAll("./static", 0777)

	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to create photo dir: %s", err.Error()))
	}

	// Static files
	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Handlers
	http.HandleFunc("/{$}", handlers.Index)
	http.HandleFunc("GET /about", handlers.About)
	http.HandleFunc("GET /upload", handlers.ShowUpload)
	http.HandleFunc("POST /upload",
		middleware.PasswordProtect(pwHash, handlers.HandleUpload))
	http.HandleFunc("GET /days/{date}", handlers.ShowDay)

	log.Fatal(http.ListenAndServe(port, nil))
}
