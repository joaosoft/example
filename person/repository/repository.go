package repository

import (
	"context"
	"github.com/joaosoft/dbr"
	errorCodes "github.com/joaosoft/example/domain/error"
	"github.com/joaosoft/example/person/domain"
	"github.com/joaosoft/logger"
)

func NewRepository(db *dbr.Dbr, log logger.ILogger) domain.IRepository {
	return &Repository{
		db:  db,
		log: log,
	}
}

func (m *Repository) GetPersonById(ctx context.Context, id int) (person *domain.Person, err error) {
	person = &domain.Person{}

	stmt := m.db.Select("id_person", "name", "age").
		From("person").
		Where("id_person = ?", id)

	var count int
	if count, err = stmt.Load(person); err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errorCodes.ErrorNotFound()
	}

	return person, nil
}

func (m *Repository) SavePerson(ctx context.Context, person *domain.SavePerson) (created *domain.CreatedPerson, err error) {
	created = &domain.CreatedPerson{}
	stmt := m.db.Insert().
		Columns("name", "age").
		Into("person").
		Return("id_person").
		Record(person)

	if _, err = stmt.Load(&created.Id); err != nil {
		return nil, err
	}

	return created, nil
}
