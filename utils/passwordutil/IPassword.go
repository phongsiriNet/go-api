package passwordutil

type IPassword interface {
	HashPassword(string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
