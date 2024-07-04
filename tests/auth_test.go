package utils_test

import (
	"testing"

	"github.com/DarkPhoenix42/LibraLynx/utils"
)

type StringBoolTest struct {
	input    string
	expected bool
}

var usernameTests = []StringBoolTest{
	{"user1", true},
	{"us", false},                       // Too short
	{"thisusernameiswaytoolong", false}, // Too long
	{"user_name", true},
	{"User123", true},
	{"user-name", false}, // Invalid character
	{"user name", false}, // Space is not allowed
	{"", false},          // Empty string
}

func TestCheckUsername(t *testing.T) {
	for _, test := range usernameTests {
		output := utils.CheckUsername(test.input)
		if output != test.expected {
			t.Errorf("Expected %v for %v, got %v", test.expected, test.input, output)
		}
	}
}

var passwordTests = []StringBoolTest{
	{"Password1!", true},
	{"password1!", false}, // No uppercase letter
	{"PASSWORD1!", false}, // No lowercase letter
	{"Password!", false},  // No digit
	{"Password1", false},  // No special character
	{"Pass1!", false},     // Too short
	{"P@ssw0rd1", true},
	{"12345678", false}, // No letters
	{"abcdefgh", false}, // No uppercase, digit, or special character
	{"A1!", false},      // Too short
	{"aB3!pass", true},
	{"aB3!PassW0rd", true},
	{"aB!7", false},
}

func TestCheckPassword(t *testing.T) {
	for _, test := range passwordTests {
		output := utils.CheckPassword(test.input)
		if output != test.expected {
			t.Errorf("Expected %v for %v, got %v", test.expected, test.input, output)
		}
	}
}

var emailTests = []StringBoolTest{
	{"test@example.com", true},
	{"user.name+tag+sorting@example.com", true},
	{"user_name@example.com", true},
	{"user@example.co.in", true},
	{"user@sub.example.com", true},
	{"user@sub-example.com", true},
	{"user@example", false},      // No TLD
	{"user@.com", false},         // No domain
	{"@example.com", false},      // No local part
	{"user@com", false},          // TLD too short
	{"user@.com.", false},        // Leading and trailing dots
	{"user@exam_ple.com", false}, // Underscore in domain
	{"user@exam!ple.com", false}, // Invalid character in domain
	{"user@exam#ple.com", false},
}

func TestCheckEmail(t *testing.T) {
	for _, test := range emailTests {
		output := utils.CheckEmail(test.input)
		if output != test.expected {
			t.Errorf("Expected %v for %v, got %v", test.expected, test.input, output)
		}
	}
}
