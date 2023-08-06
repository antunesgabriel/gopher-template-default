package helper

type PasswordHelper interface {
	Hash(password string) ([]byte, error)
	Compare(password, hashedPassword string) error
}
