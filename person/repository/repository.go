package repository

import (
	"context"
	"database/sql"
	"github.com/joaosoft/example/person/domain"
	"math/rand"
)

func NewRepository(db *sql.DB) domain.IRepository {
	return &Repository{
		db: db,
	}
}

func (m *Repository) GetPersonById(ctx context.Context, id int) (*domain.Person, error) {
	return &domain.Person{
		Id:   id,
		Name: domain.GenerateName(10),
		Age:  rand.Intn(100),
	}, nil
}
