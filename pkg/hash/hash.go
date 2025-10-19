package hash

import "golang.org/x/crypto/bcrypt"

func GenHash(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 14)
	return string(bytes), err
}

func EqToHash(s, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(s)) == nil
}
