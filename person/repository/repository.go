package repository

import (
	"context"
	"database/sql"
	"github.com/joaosoft/example/domain/person"
	"math/rand"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) person.IRepository {
	return &Repository{
		db: db,
	}
}

func (m *Repository) GetPersonById(ctx context.Context, id int) (*person.Person, error) {
	return &person.Person{
		Id:   id,
		Name: person.RandomString(10),
		Age:  rand.Intn(100),
	}, nil
}
