package util

import "strings"

//PasswordIsValid checks if a given password meets the password security requirements
func PasswordIsValid(password string) bool {
	return strings.ContainsAny(strings.ToLower(password), "abcdefghijklmnopqrstuvwxyz") && strings.ContainsAny(password, "0123456789") && len(password) >= 8
}
