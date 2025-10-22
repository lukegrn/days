package db

import (
	"database/sql"
	"fmt"

	"github.com/lukegrn/days/pkg/img"
	_ "github.com/mattn/go-sqlite3"
)

type Inst struct {
	sql.DB
}

type ImgOverview struct {
	PrevDate string
	Cur      *img.Img
	NextDate string
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

func (i *Inst) GetImage(date string) (ImgOverview, error) {
	imgOverview := ImgOverview{Cur: &img.Img{}}
	row := i.QueryRow(`SELECT * FROM
			(SELECT
				IFNULL(LAG(date) OVER (ORDER BY date), '') as prev_date,
				date as cur_date,
				IFNULL (LEAD(date) OVER (ORDER BY date), '') as next_date,
				path,
				caption
			FROM images)
		WHERE cur_date= ? ;`, date)

	err := row.Scan(&imgOverview.PrevDate,
		&imgOverview.Cur.Date, &imgOverview.NextDate, &imgOverview.Cur.Path, &imgOverview.Cur.Caption)

	if err != nil {
		return ImgOverview{}, fmt.Errorf("Failed to get img: %s", err.Error())
	}

	return imgOverview, nil
}

func (i *Inst) GetAllImages() ([]img.Img, error) {
	images := make([]img.Img, 0)
	rows, err := i.Query(`SELECT date, path, caption FROM images ORDER BY date`)
	if err != nil {
		return images, fmt.Errorf("Failed to get images: %s", err.Error())
	}

	for rows.Next() {
		image := img.Img{}
		err = rows.Scan(&image.Date, &image.Path, &image.Caption)
		if err != nil {
			return images, fmt.Errorf("Failed to get images: %s", err.Error())
		}

		images = append(images, image)
	}

	return images, nil
}
