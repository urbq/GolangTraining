package main

import (
	"fmt"
	"time"

	"github.com/urbq/GolangTraining/exercises_02/6/primes"
)

//The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//Find the sum of all the primes below two million.
func main() {
	start := time.Now()
	primes := primes.Test2(10)
	//primes := primes.Test(100000)
	//primes := primes.GetPrimes(100000)
	fmt.Printf("Primes\t\t%v\n", primes)
	var sum int
	for _, p := range primes {
		sum += p
	}
	fmt.Printf("Sum primes\t%v\n", sum)

	elapsed := time.Since(start)
	fmt.Printf("Primes took %s", elapsed)
}
