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
	err := db.Get(trivia, "select * from trivia offset random()*(select max(trivia_id) from trivia) limit 1")
	if err != nil {
		return nil, err
	}
	return trivia, nil
}

func (db *DB) GetTrivia(id int64) (*Trivia, error) {
	trivia := &Trivia{}
	err := db.Get(trivia, "select * from trivia where trivia_id = $1", id)
	if err != nil {
		return nil, err
	}
	return trivia, nil
}

func (db *DB) GetAllTrivia() ([]*Trivia, error) {
	var trivia []*Trivia
	err := db.Select(&trivia, "select * from trivia limit 10")
	if err != nil {
		return nil, err
	}
	return trivia, nil
}
