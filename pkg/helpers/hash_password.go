package helpers

import "golang.org/x/crypto/bcrypt"

// HashPassword securely hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	// Use bcrypt with default cost (12)
	// Higher cost means more computational time but more secure
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash compares a plain text password with its hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
