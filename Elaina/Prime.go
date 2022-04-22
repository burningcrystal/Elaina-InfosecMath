package Elaina

func IsPrime(p int) bool {
	//start := time.Now() // 获取当前时间
	if p < 2 {
		return false
	}
	//SqrtN := int(math.Sqrt(float64(a)))
	for i := 2; i <= p/i; i++ {
		if p%i == 0 {
			return false
		}
	}
	//elapsed := time.Since(start)
	//log.Println("该函数执行完成耗时：", elapsed)
	return true
}

func DevidePrime(p int) map[int]int { //底数是key，幂数是value
	res := make(map[int]int)
	for i := 2; i <= p/i; i++ {
		if p%i == 0 {
			s := 0
			for p%i == 0 {
				p /= i
				s++
			}
			res[i] = s
		}
	}
	if p > 1 {
		res[p] = 1
	}
	return res
}

func GetPrime(n int) []int {
	prime := make([]int, n+1)
	sieve := make([]bool, n+1) //真值表示是合数，假值表示是质数。
	var counter int = 0

	for i := 2; i <= n; i++ {
		if !sieve[i] { //判断是不是质数
			prime[counter] = i
			counter = counter + 1
		}
		for j := 0; prime[j]*i <= n; j = j + 1 {
			sieve[i*prime[j]] = true
			if i%prime[j] == 0 {
				break
			}
		}
	}
	return prime[:counter]
}

func Phi(n int) int {
	normalization := DevidePrime(n)
	res := n
	for base, _ := range normalization {
		res = res - (res / base)
	}
	return res
}

func BatchPhi(n int) []int {
	primes := make([]int, n+1)
	sieve := make([]bool, n+1)
	euler := make([]int, n+1)
	var counter int = 0

	euler[0] = 114514
	euler[1] = 1
	for i := 2; i <= n; i++ {
		if !sieve[i] {
			primes[counter] = i
			counter++
			euler[i] = i - 1
		}

		for j := 0; primes[j]*i <= n; j++ {
			tmp := primes[j] * i
			sieve[tmp] = true

			if i%primes[j] == 0 {
				euler[tmp] = euler[i] * primes[j]
				break
			}
			euler[tmp] = euler[i] * (primes[j] - 1)
		}
	}
	return euler
}
