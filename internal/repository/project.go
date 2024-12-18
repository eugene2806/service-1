package repository

import (
	"database/sql"
	"service-1/storage"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(storage *storage.Storage) *ProjectRepository {
	return &ProjectRepository{
		db: storage.GetDB(),
	}
}
