package repositories

import (
	"github.com/TusharChoudharykp/Email-Campaign-Tool/config"
	models "github.com/TusharChoudharykp/Email-Campaign-Tool/models"
)

func CreateUser(user models.User) error {
	query := `
	INSERT INTO users(name, email, password, role)
	VALUES (?, ?, ?, ?)
	`

	_, err := config.DB.Exec(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
	)

	return err
}

func GetUserByEmail(email string) (*models.User, error) {
	query := `
	SELECT id, name, email, password, role, created_at
	FROM users
	WHERE email = ?
	`

	var user models.User

	err := config.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
