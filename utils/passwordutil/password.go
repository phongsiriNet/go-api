package passwordutil

import "golang.org/x/crypto/bcrypt"

type passwordUtil struct{}

func NewPasswordUtil() IPassword {
	return &passwordUtil{}
}

func (p *passwordUtil) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err

}

func (p *passwordUtil) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
