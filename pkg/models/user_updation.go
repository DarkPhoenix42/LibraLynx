package models

import (
	"github.com/DarkPhoenix42/LibraLynx/db"
	"github.com/DarkPhoenix42/LibraLynx/pkg/utils"
)

func AddUser(username, email, password string, is_admin bool) error {
	query := "INSERT INTO users (username, email, password, is_admin) VALUES (?, ?, ?, ?)"
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	_, err = db.DB.Exec(query, username, email, hashedPassword, is_admin)
	if err != nil {
		return err
	}

	return nil
}

func MakeUserAdmin(user_id int) error {
	query := "UPDATE users SET is_admin = 1 WHERE user_id = ?"
	_, err := db.DB.Exec(query, user_id)
	return err
}

func GetUserAdminRequestStatus(user_id int) (string, error) {
	query := "SELECT admin_request_status FROM users WHERE user_id = ?"
	var status string
	err := db.DB.QueryRow(query, user_id).Scan(&status)
	return status, err
}

func UpdateUserAdminRequestStatus(user_id int, status string) error {
	query := "UPDATE users SET admin_request_status = ? where user_id = ?"
	_, err := db.DB.Exec(query, status, user_id)
	return err
}
