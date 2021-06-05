package hashhelper

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GetHashedValue(PWD string) ([]byte, error) {
	password := []byte(PWD)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return []byte(``), fmt.Errorf("Generate pwd with bcrypt bumped into error: %v", err)
	}
	return hashedPassword, nil
}
