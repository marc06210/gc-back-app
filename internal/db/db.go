package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marc06210/gc-back-app/internal/model"
)

type DB struct {
	pool *pgxpool.Pool
}

func New(username, password, dbname, host string, port int) (*DB, error) {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", username, password, host, port, dbname)
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to b: %w", err)
	}
	return &DB{pool: pool}, nil
}

// returns all the publications
func (db *DB) GetAllPublications() ([]model.Publication, error) {
	// TODO: can I implement a scanner to avoid this full arguments SELECT
	query := "SELECT id, title, description, icon, host, url, creationts FROM publication"
	rows, err := db.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	publications := make([]model.Publication, 0)
	for rows.Next() {
		var publication model.Publication
		var host sql.NullString
		var url sql.NullString
		var icon model.Icon
		err := rows.Scan(&publication.Id, &publication.Title, &publication.Description, &icon, &host, &url, &publication.Creationts)
		publication.Host = db.string(host)
		publication.Url = db.string(url)
		publication.Icon = icon.String()
		if err != nil {
			return nil, err
		}
		publications = append(publications, publication)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return publications, nil
}

// converts a sqlNullString into a string
func (db *DB) string(nullString sql.NullString) string {
	if nullString.Valid {
		return nullString.String
	}
	return ""
}

func (db *DB) Close() {
	db.pool.Close()
}
