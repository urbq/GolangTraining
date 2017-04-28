package main

import "fmt"
import "github.com/urbq/GolangTraining/exercise2/6/primes"

func main() {
	numbers := []int{2, 3, 4, 5, 6, 7, 213132}
	fmt.Printf("Numbers\t\t%v\n", numbers)
	primes := primes.GetPrimes(numbers...)
	fmt.Printf("Primes\t\t%v\n", primes)
	var sum int
	for _, p := range primes {
		sum += p
	}
	fmt.Printf("Sum primes\t%v\n", sum)
}
