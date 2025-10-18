package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Inst struct {
	sql.DB
}

type img struct {
	Date    string
	Caption string
	Path    string
}

var i *Inst = &Inst{}

func Get() *Inst {
	return i
}

func (i *Inst) SetupDB(f string) error {
	db, err := sql.Open("sqlite3", f)
	i.DB = *db
	if err != nil {
		return fmt.Errorf("Failed to set up db: %s", err.Error())
	}

	create_tbl := `
	CREATE TABLE IF NOT exists
		images
		(
			date TEXT PRIMARY KEY NOT NULL,
			path TEXT NOT NULL,
			caption TEXT NOT NULL
		)
`
	_, err = db.Exec(create_tbl)
	if err != nil {
		return fmt.Errorf("Failed to set up db: %s", err.Error())
	}

	return nil
}

func (i *Inst) PutImage(date, path, caption string) error {
	trx, err := i.Begin()
	defer trx.Rollback()
	if err != nil {
		return fmt.Errorf("Failed to begin transaction: %s", err.Error())
	}

	stmt, err := trx.Prepare(`INSERT INTO images (date, path, caption) VALUES (?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("Failed to prepare transaction: %s", err.Error())
	}

	_, err = stmt.Exec(date, path, caption)
	if err != nil {
		return fmt.Errorf("Failed to execute transaction: %s", err.Error())
	}

	err = trx.Commit()
	if err != nil {
		return fmt.Errorf("Failed to commit transaction: %s", err.Error())
	}

	return nil
}

func (i *Inst) GetImage(date string) (img, error) {
	var id string
	var path string
	var caption string
	row := i.QueryRow(`SELECT date, path, caption FROM images WHERE date = ?`, date)

	err := row.Scan(&id, &path, &caption)
	if err != nil {
		return img{}, fmt.Errorf("Failed to get img: %s", err.Error())
	}

	return img{Date: id, Path: path, Caption: caption}, nil
}
