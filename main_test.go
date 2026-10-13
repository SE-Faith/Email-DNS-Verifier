package main

import (
	"testing"
	"strings"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"omo@gmail.com", true},
		{"omogmail.com", false},
		{"@gmail.com", false},
		{"omo@", false},
		{"omo@@gmail.com", false},
		{"", false},
		{"omo@gmail..com", false},
		{"omo@.gmail.com", false},
		{"omo@gmail.com.", false},
		{"omo @gmail.com", false},
		
	}
	for _, test := range tests {
		result := isValidEmail(test.email)
		if result != test.expected {
			t.Errorf("isValidEmail(%q) = %v, want %v", test.email, result, test.expected)
		}
		if result {
			parts := strings.Split(test.email, "@")
			domain := parts[1]
			checkIP(domain)
			checkMX(domain)
			checkSPF(domain)
			checkDMARC(domain)
		}
	}
}


