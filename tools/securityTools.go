package tools

import "golang.org/x/crypto/bcrypt"

//HashPassword generate hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
//CheckHashAndPassword check password
func CheckHashAndPassword(hash,password  string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}