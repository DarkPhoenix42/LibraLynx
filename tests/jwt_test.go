package utils_test

import (
	"testing"

	"github.com/DarkPhoenix42/LibraLynx/pkg/utils"
)

type Test struct {
	user_id  int
	is_admin bool
}

var jwtTests = []Test{
	{1, true},
	{2, false},
	{3, true},
}

func TestCreateToken(t *testing.T) {
	for _, test := range jwtTests {
		token, err := utils.CreateToken(test.user_id, test.is_admin)
		if err != nil {
			t.Errorf("Error creating token: %v", err)
		}
		user_id, is_admin, err := utils.DecodeToken(token)
		if err != nil {
			t.Errorf("Error decoding token: %v", err)
		}
		if user_id != test.user_id || is_admin != test.is_admin {
			t.Errorf("Expected %v, %v, got %v, %v", test.user_id, test.is_admin, user_id, is_admin)
		}
	}
}

func TestDecodeToken(t *testing.T) {
	for _, test := range jwtTests {
		token, err := utils.CreateToken(test.user_id, test.is_admin)
		if err != nil {
			t.Errorf("Error creating token: %v", err)
		}
		user_id, is_admin, err := utils.DecodeToken(token)
		if err != nil {
			t.Errorf("Error decoding token: %v", err)
		}
		if user_id != test.user_id || is_admin != test.is_admin {
			t.Errorf("Expected %v, %v, got %v, %v", test.user_id, test.is_admin, user_id, is_admin)
		}
	}
}
