package services

import (
	"database/sql"
	"job_portal/internal/models"
	"job_portal/internal/repository"
)

func GetUserById(db *sql.DB, id int) (*models.User, error) {
	return repository.GetUserById(db, id)
}

