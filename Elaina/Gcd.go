package Elaina

func Gcd(a int, b int) int {
	var c int //c是余数
	for b != 0 {
		c = a % b
		a = b
		b = c
	}
	return a
}

//todo 使用迭代法重写
func ExGcd(a, b int, x *int, y *int) int {
	if b == 0 {
		*x = 1
		*y = 0
		return a
	}

	d := ExGcd(b, a%b, y, x)
	(*y) -= (a / b) * (*x)
	return d
}

func Lcm(a, b int) int {
	return (a * b) / Gcd(a, b)
}
