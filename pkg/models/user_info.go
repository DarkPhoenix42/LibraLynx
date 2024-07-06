package models

import (
	"github.com/DarkPhoenix42/LibraLynx/db"
	"github.com/DarkPhoenix42/LibraLynx/pkg/types"
)

func CheckFirstUser() (bool, error) {
	query := "SELECT COUNT(*) FROM users"
	row := db.DB.QueryRow(query)

	var count int
	err := row.Scan(&count)
	if err != nil || count != 0 {
		return false, err
	}

	return true, nil

}

func GetAdminRequests() ([]types.AdminRequest, error) {
	query := "SELECT user_id, username FROM users WHERE admin_request_status = 'pending'"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var requests []types.AdminRequest
	for rows.Next() {
		var request types.AdminRequest
		err := rows.Scan(&request.UserID, &request.Username)
		if err != nil {
			return nil, err
		}

		requests = append(requests, request)
	}

	return requests, nil
}

func GetUserByUsername(username string) (*types.User, error) {
	query := "SELECT * FROM users WHERE username = ?"
	row := db.DB.QueryRow(query, username)

	var user types.User
	err := row.Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.IsAdmin,
		&user.AdminRequestStatus)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
