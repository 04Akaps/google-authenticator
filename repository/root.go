package repository

import (
	"database/sql"
	"strings"

	"demo-scrapping/config"
	"demo-scrapping/types/schema"

	_ "github.com/go-sql-driver/mysql"
)

type repository struct {
	cfg *config.Config
	db  *sql.DB
}

type RepositoryImpl interface {
	Add(url, cardSelector, innerSelector string, tag []string) error
	Update(url, cardSelector, innerSelector string, tag []string) error
	Delete(url string) error
	View(url string) (*schema.Admin, error)
	ViewAll() ([]*schema.Admin, error)
}

const (
	admin = "Scrapping.Admin"
)

func NewRepositry(cfg *config.Config) (RepositoryImpl, error) {
	r := &repository{cfg: cfg}

	dbCfg := cfg.DB
	var err error

	if r.db, err = sql.Open(dbCfg.Database, dbCfg.URL); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

func (r *repository) Add(url, cardSelector, innerSelector string, tag []string) error {
	_, err := r.db.Exec(
		"INSERT INTO Scrapping.Admin(url, tag, cardSelector, innerSelector) VALUES(?, ?, ?, ?)",
		url, strings.Join(tag, " "), cardSelector, innerSelector,
	)

	return err
}

func (r *repository) View(url string) (*schema.Admin, error) {
	s := new(schema.Admin)
	qs := query([]string{"SELECT * FROM", admin, "WHERE url = ?"})

	err := r.db.QueryRow(qs, url).Scan(
		&s.ID,
		&s.URL,
		&s.CardSelector,
		&s.InnerSelector,
		&s.Tag,
		&s.CreatedAt,
		&s.UpdatedAt,
	)

	return s, err
}

func (r *repository) ViewAll() ([]*schema.Admin, error) {
	qs := query([]string{"SELECT * FROM", admin})

	if cursor, err := r.db.Query(qs); err != nil {
		return nil, err
	} else {
		defer cursor.Close()

		var result []*schema.Admin

		for cursor.Next() {
			s := new(schema.Admin)

			if err = cursor.Scan(
				&s.ID,
				&s.URL,
				&s.CardSelector,
				&s.InnerSelector,
				&s.Tag,
				&s.CreatedAt,
				&s.UpdatedAt,
			); err != nil {
				return nil, err
			} else {
				result = append(result, s)
			}
		}

		if len(result) == 0 {
			return []*schema.Admin{}, nil
		} else {
			return result, nil
		}
	}
}

func (r *repository) Update(url, cardSelector, innerSelector string, tag []string) error {
	q := query([]string{"UPDATE", admin, "SET tag = ?, cardSelector = ?, innerSelector = ? WHERE url = ?"})
	_, err := r.db.Exec(q, strings.Join(tag, " "), cardSelector, innerSelector, url)
	return err
}

func (r *repository) Delete(url string) error {
	q := query([]string{"DELETE FROM", admin, "WHERE url = ?"})
	_, err := r.db.Exec(q, url)
	return err
}

func query(qs []string) string {
	return strings.Join(qs, " ") + ";"
}
