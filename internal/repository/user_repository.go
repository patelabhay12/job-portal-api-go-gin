package repository

import (
	"database/sql"
	"job_portal/internal/models"
)

func CreateUser(db *sql.DB, user *models.User) error {
	_, err := db.Exec(`INSERT INTO users (username,password,email) VALUES (?,?,?)`, user.Username, user.Password, user.Email)

	return err
}

func GetUserById(db *sql.DB, id int) (*models.User, error) {
	var user models.User
	var profilePicture sql.NullString

	err := db.QueryRow(`
        SELECT id, username, email, profile_picture, created_at, updated_at
        FROM users WHERE id = ?
    `, id).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&profilePicture,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	if profilePicture.Valid {
		user.ProfilePicture = profilePicture.String
	} else {
		user.ProfilePicture = ""
	}

	return &user, nil
}
func LoginUser(db *sql.DB, username string, password string) (*models.User, error) {
    user := &models.User{}

    // Scan directly into proper types, handle NULL profile_picture
    err := db.QueryRow(`
        SELECT id, username, password, email, created_at, updated_at, is_admin,
               COALESCE(profile_picture, '') AS profile_picture
        FROM users
        WHERE username = ?
    `, username).Scan(
        &user.Id,
        &user.Username,
        &user.Password,
        &user.Email,
        &user.CreatedAt,
        &user.UpdatedAt,
        &user.IsAdmin,
        &user.ProfilePicture,
    )

    if err != nil {
        return nil, err
    }

    return user, nil
}
