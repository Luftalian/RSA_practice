package crypto

type Crypto interface {
	Encrypt(a int) (int, error)
	Decrypt(b int) (int, error)
}
