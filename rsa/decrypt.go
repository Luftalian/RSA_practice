package rsa

import "fmt"

func (r RSA) Decrypt(b int) (int, error) {
	// d: 暗号文
	// n, d: 秘密鍵
	if b >= r.n {
		return 0, fmt.Errorf("c must be less than n")
	}
	return pow(b, r.d, r.n), nil
}
