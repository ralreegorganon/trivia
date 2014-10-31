package trivia

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
}

func (db *DB) Open() error {
	d, err := sqlx.Open("postgres", "timezone=UTC sslmode=disable")
	if err != nil {
		return err
	}

	db.DB = d

	return nil
}

func (db *DB) GetRandomTrivia() (*Trivia, error) {
	trivia := &Trivia{}
	err := db.Get(trivia, "select question, answer from trivia offset random()*(select max(trivia_id) from trivia) limit 1")
	if err != nil {
		return nil, err
	}
	return trivia, nil
}
