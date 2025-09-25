package services

import (
	"database/sql"
	"job_portal/internal/models"
	"job_portal/internal/repository"
	"job_portal/internal/utiles"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *sql.DB, user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return repository.CreateUser(db, user)
}

func LoginUser(db *sql.DB, username string, password string) (string, error) {
	user, err := repository.LoginUser(db, username, password)

	if err != nil {
		return "check this", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "in this", err
	}

	return utiles.GenerateJWT(username, user.Id, user.IsAdmin)
}
