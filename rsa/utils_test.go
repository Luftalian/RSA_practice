package rsa

import (
	"testing"
)

func Test_pow(t *testing.T) {
	type args struct {
		a int
		e int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Case 1", args{2, 3, 5}, 3},     // 2^3 mod 5 = 3
		{"Case 2", args{3, 4, 7}, 4},     // 3^4 mod 7 = 4
		{"Case 3", args{5, 0, 11}, 1},    // 5^0 mod 11 = 1
		{"Case 4", args{2, 10, 13}, 10},  // 2^10 mod 13 = 10
		{"Case 5", args{2, 11, 13}, 7},   // 2^11 mod 13 = 7
		{"Case 6", args{100, 0, 100}, 1}, // 100^0 mod 100 = 1
		{"Case 7", args{100, 1, 100}, 0}, // 100^1 mod 100 = 0
		{"Case 8", args{0, 100, 100}, 0}, // 0^100 mod 100 = 0
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow(tt.args.a, tt.args.e, tt.args.n); got != tt.want {
				t.Errorf("pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Case 1", args{3, 11}, 1},
		{"Case 2", args{7, 13}, 1},
		{"Case 3", args{5, 17}, 1},
		{"Case 4", args{7, 17}, 1},
		{"Case 5", args{3, 6}, 3},
		{"Case 6", args{6, 3}, 3},
		{"Case 7", args{6, 6}, 6},
		{"Case 8", args{6, 4}, 2},
		{"Case 9", args{14, 16}, 2},
		{"Case 10", args{47, 74}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_modInv(t *testing.T) {
	type args struct {
		a int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Case 1", args{3, 11}, 4}, // 3の逆数は 4 (mod 11)
		{"Case 2", args{7, 13}, 2}, // 7の逆数は 2 (mod 13)
		{"Case 3", args{5, 17}, 7}, // 5の逆数は 7 (mod 17)
		{"Case 4", args{7, 17}, 5}, // 7の逆数は 5 (mod 17)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modInv(tt.args.a, tt.args.m); got != tt.want {
				t.Errorf("modInv() = %v, want %v", got, tt.want)
			}
		})
	}
}
