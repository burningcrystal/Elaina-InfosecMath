package Elaina

import (
	"errors"
	"log"
)

func Qmi(a, k, p int) int {
	res := 1
	for k != 0 {
		if k&1 != 0 {
			res = res * a % p
		}
		k >>= 1
		a = a * a % p
	}
	return res
}

func GetInverse(a, b, m int) (int, error) { //b是要求的逆元的对象
	if IsPrime(m) {
		if m%b == 0 {
			return 114514, errors.New("the there is no Inverse element about a/b congruent with m!")
		}
		return Pow(b, (m - 1)), nil
	}
	b_1, _, err := LinearCongruenceEquation(b, 1, m)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return b_1, nil
}

func LinearCongruenceEquation(a, b, m int) (int, int, error) { //ax同余于b(mod m)
	x, y := 0, 0
	d := ExGcd(a, m, &x, &y)
	if d%b == 0 {
		return x, y * (-1), nil
	}
	return 114514, 114514, errors.New("the there is no solution about this Linear congruence equation!")
}
