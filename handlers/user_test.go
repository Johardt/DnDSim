package handlers

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test", false},
		{"test@", false},
		{"test@domain", false},
		{"test@domain.", false},
		{"@.com", false},
		{"test@com", false},
		{"test@domain.com", true},
	}

	for _, test := range tests {
		if isValidEmail(test.email) != test.expected {
			t.Errorf("Expected %v for %s, but got %v", test.expected, test.email, !test.expected)
		}
	}
}

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"", false},
		{"123456789101112131415161718192021222324252627282930", true},
	}

	for _, test := range tests {
		if isValidPassword(test.password) != test.expected {
			t.Errorf("Expected %v for %s, but got %v", test.expected, test.password, !test.expected)
		}
	}
}

func TestHashPassword(t *testing.T) {
	tests := []struct {
		password string
	}{
		{"test"},
		{"password"},
	}
	for _, test := range tests {
		hashed, err := HashPassword(test.password)
		if err != nil {
			t.Errorf("Error hashing password: %v", err)
		}
		if len(hashed) == 0 {
			t.Errorf("Expected hashed password, but got empty string")
		}
		hashed2, _ := HashPassword(test.password)
		if hashed == hashed2 {
			t.Errorf("Expected different hashes for different passwords, but got the same hash")
		}
		hashed3, _ := HashPassword("test password")
		if hashed == hashed3 {
			t.Errorf("Expected different hashes for different passwords, but got the same hash")
		}
	}
}

func TestVerifyPassword(t *testing.T) {
	tests := []struct {
		password string
	}{
		{"test"},
	}
	for _, test := range tests {
		hashed, err := HashPassword(test.password)
		if err != nil {
			t.Errorf("Error hashing password: %v", err)
		}
		if VerifyPassword(test.password, hashed) != nil {
			t.Errorf("Expected password to match hash, but it didn't")
		}
		if VerifyPassword("wrong password", hashed) == nil {
			t.Errorf("Expected password not to match hash, but it did")
		}
	}
}
