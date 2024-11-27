package handlers

import "testing"

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
