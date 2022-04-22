package Elaina

func GetDigits(n int) int { //get the digits of the number n
	if n == 0 {
		return 1
	}
	digits := 0
	for n != 0 {
		n /= 10
		digits++
	}
	return digits
}

func Pow(base, power int) int {
	for i := 0; i < power; i++ {
		base *= base
	}
	return base
}
