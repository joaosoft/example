package repository

import (
	"database/sql"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	db *sql.DB
}

type RepositoryMock struct {
	mock.Mock
}
