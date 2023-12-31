package store

import( 
"database/sql"

_"github.com/lib/pq"
)

type Store struct {
	config *Config
	db     *sql.DB
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.databaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db


	return nil
}

func (s *Store) Close() {
	s.db.Close()
}
