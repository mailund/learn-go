package main

import "fmt"

func egcd(a, b int64) (int64, int64, int64) {
	oldR, r := a, b
	oldS, s := int64(1), int64(0)
	oldT, t := int64(0), int64(1)
	for r != 0 {
		quotient := oldR / r
		oldR, r = r, oldR-quotient*r
		oldS, s = s, oldS-quotient*s
		oldT, t = t, oldT-quotient*t
	}
	return oldR, oldS, oldT
}

func multmod(a, b, m int64) int64 {
	var result int64 = 0
	a %= m
	b %= m

	for b != 0 {
		if b%2 != 0 {
			result = (result + a) % m
		}

		a = (a * 2) % m
		b /= 2
	}

	return result
}

func mod(a, m int64) int64 {
	a %= m
	if a < 0 {
		a += m
	}
	return a
}

func crt(a, n, b, m int64) (int64, int64) {
	K := n * m
	_, s, t := egcd(n, m)
	atm := multmod(multmod(a, mod(t, K), K), m, K)
	bsn := multmod(multmod(b, mod(s, K), K), n, K)
	x := mod(atm+bsn, K)
	return x, K
}

func main() {
	var T int64
	fmt.Scan(&T)
	for ; T > 0; T-- {
		var a, n, b, m int64
		fmt.Scan(&a, &n, &b, &m)
		x, K := crt(a, n, b, m)
		fmt.Println(x, K)
	}
}
