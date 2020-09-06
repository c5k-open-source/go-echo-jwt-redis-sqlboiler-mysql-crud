package security

import (
	"fmt"

	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	password = fmt.Sprintf("%s%s", config.Config.PasswordHashSalt, password)
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	fmt.Println(hashedPassword)
	fmt.Println(password)
	password = fmt.Sprintf("%s%s", config.Config.PasswordHashSalt, password)
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
