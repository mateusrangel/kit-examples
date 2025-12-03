package repository

import (
	"database/sql"

	"github.com/mateusrangel/kit/fsm/internal/domain"
	_ "modernc.org/sqlite"
)

type RepoImplSqlite struct {
	db *sql.DB
}

func New() (*RepoImplSqlite, error) {
	db, err := sql.Open("sqlite", "disputes.sqlite")
	if err != nil {
		return nil, err
	}
	query := `
  	CREATE TABLE IF NOT EXISTS disputes (
		id TEXT PRIMARY KEY,
  		state TEXT NOT NULL
	);
  `
	if _, err := db.Exec(query, nil); err != nil {
		return nil, err
	}
	return &RepoImplSqlite{db}, nil
}

func (r *RepoImplSqlite) CreateDispute(d *domain.Dispute) error {
	if _, err := r.db.Query("INSERT INTO disputes VALUES (?, ?)", d.Id, d.State); err != nil {
		return err
	}
	return nil
}

func (r *RepoImplSqlite) UpdateState(id, newState string) error {
	if _, err := r.db.Query("UPDATE disputes SET state=? where id=?", newState, id); err != nil {
		return err
	}
	return nil
}
