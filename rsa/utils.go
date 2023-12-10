package rsa

type RSA struct {
	n int // 公開鍵
	e int // 公開鍵
	d int // 秘密鍵
}

// p, q: 素数
func NewRSA(p, q int) RSA {
	n := p * q
	phi := (p - 1) * (q - 1)
	e := 3
	for gcd(phi, e) != 1 {
		e += 2
	}
	d := modInv(e, phi)
	return RSA{n, e, d}
}

// a^e mod n
func pow(a, e, n int) int {
	if e == 0 {
		return 1
	}
	if e%2 == 0 {
		return pow((a*a)%n, e/2, n)
	}
	return (a * pow(a, e-1, n)) % n
}

// 最大公約数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// モジュラ逆数
func modInv(a, m int) int {
	m0, x0, x1 := m, 0, 1
	for a > 1 {
		q := a / m
		m, a = a%m, m
		x0, x1 = x1-q*x0, x0
	}
	// 結果が負の場合、正の範囲に調整
	if x1 < 0 {
		x1 += m0
	}
	return x1
}
