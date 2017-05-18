package primes

import (
	"math"
)

//GetPrimes my first idea
func GetPrimes(number int) []int {
	//create slice with all numbers
	var numbers []int
	//starting with 2
	numbers = append(numbers, 2)
	//and then only even
	i := 3
	for i <= number {
		numbers = append(numbers, i)
		i += 2
	}
	return getPrimesForSlice(numbers...)
}

func getPrimesForSlice(numbers ...int) []int {
	//create slice with all primes from given numbers
	var result []int
	for _, n := range numbers {
		if isPrime(n) {
			result = append(result, n)
		}
	}
	return result
}

func isPrime(s int) bool {
	if s == 1 {
		return false
	}
	for i := 2; i < s; i++ {
		if s%i == 0 {
			return false
		}
	}
	return true
}

//Primitive implementation
func Primitive(treshold int) []int {
	primes := []int{2}

	p := 3
	for p < treshold {
		sqrt := int(math.Sqrt(float64(p)))
		isPrime := true
		divider := 3
		for divider <= sqrt {
			if p%divider == 0 {
				isPrime = false
			}
			divider += 2
		}
		if isPrime {
			primes = append(primes, p)
		}
		p += 2
	}
	return primes
}

//Erastosthenes filter implementation
func Erastosthenes(treshold int) int {
	sqrt := int(math.Sqrt(float64(treshold)))
	// fmt.Println(sqrt)
	filter := make([]bool, treshold)
	//fmt.Println(filter)
	{
		n := 4
		for n < treshold {
			filter[n] = true
			n += 2
		}
	}
	//fmt.Println(filter)
	{
		n := 3
		for n <= sqrt {
			if filter[n] == false {
				m := n * n
				for m <= treshold {
					filter[m] = true
					m += 2 * n
				}
			}
			n += 2
		}
	}
	//fmt.Println(filter)
	sum := 2
	m := 3
	for m <= treshold {
		if filter[m] == false {
			sum += m
		}
		m += 2
	}
	return sum
}
