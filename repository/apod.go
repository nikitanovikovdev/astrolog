package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nikitanovikovdev/astrolog/errorutil"
	"github.com/nikitanovikovdev/astrolog/repository/entity"
)

type APODRepo struct {
	db        *sqlx.DB
	tableName string
}

func NewAPODRepo(db *sqlx.DB) *APODRepo {
	return &APODRepo{
		db:        db,
		tableName: "content",
	}
}

func (r APODRepo) ContentByDate(date string) (entity.APODContent, error) {
	var content entity.APODContent
	query := fmt.Sprintf(`SELECT * FROM %s WHERE date=$1`, r.tableName)
	err := r.db.Get(&content, query, date)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.APODContent{}, errorutil.ErrNoRecords
		}
		return entity.APODContent{}, err
	}

	return content, nil
}

func (r APODRepo) AllContent() ([]entity.APODContent, error) {
	var contents []entity.APODContent
	query := fmt.Sprintf(`SELECT * FROM %s`, r.tableName)
	if err := r.db.Select(&contents, query); err != nil {
		return []entity.APODContent{}, err
	}

	return contents, nil
}

func (r APODRepo) InsertContent(data entity.APODContent) error {
	query := fmt.Sprintf(`INSERT INTO %s (copyright, explanation, hdurl, media_type, service_version, title, url, date)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, r.tableName)
	_, err := r.db.Exec(query, data.Copyright, data.Explanation, data.HdURL, data.MediaType, data.ServiceVersion, data.Title, data.URL, data.Date)
	return err
}
