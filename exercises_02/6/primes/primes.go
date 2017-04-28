package primes

func GetPrimes(numbers ...int) []int {

	isPrime := func(s int) (int, bool) {
		if s == 1 {
			return s, false
		}
		for i := 2; i < s; i++ {
			if s%i == 0 {
				return s, false
			}
		}
		return s, true
	}

	var result []int
	for _, n := range numbers {
		value, isPrime := isPrime(n)
		if isPrime {
			result = append(result, value)
		}
	}
	return result
}
