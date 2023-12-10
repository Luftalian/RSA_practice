package rsa

import "fmt"

func (r RSA) Encrypt(a int) (int, error) {
	// a: 平文
	// n, e: 公開鍵
	if a >= r.n {
		return 0, fmt.Errorf("a must be less than n")
	}
	return pow(a, r.e, r.n), nil
}
