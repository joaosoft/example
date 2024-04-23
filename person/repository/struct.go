package repository

import (
	"github.com/joaosoft/dbr"
	"github.com/joaosoft/logger"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	db  *dbr.Dbr
	log logger.ILogger
}

type RepositoryMock struct {
	mock.Mock
}
