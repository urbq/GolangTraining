package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/urbq/GolangTraining/exercises_02/6/primes"
)

//The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//Find the sum of all the primes below two million.
func main() {
	var treshold int
	var primitive bool
	var erastostehens bool
	if len(os.Args) < 2 {
		fmt.Println("usage: main.go -treshold=number -primitive=bool -erastosthenes=bool")
		os.Exit(0)
	}
	flag.IntVar(&treshold, "treshold", 20, "the treshold")
	flag.BoolVar(&primitive, "primitive", false, "use primitive algorythm")
	flag.BoolVar(&erastostehens, "erastostehens", true, "use erastostehens algorythm")
	flag.Parse()
	if primitive {
		fmt.Println("primitive:")
		start := time.Now()
		primitive := primes.Primitive(treshold)
		//uncomment me to print all primes
		//fmt.Printf("Primes\t\t%v\n", primitive)
		var sum int
		for _, p := range primitive {
			sum += p
		}
		fmt.Printf("Sum primes\t%v\n", sum)
		elapsed := time.Since(start)
		fmt.Printf("Primes took %s\n", elapsed)
	}
	if erastostehens {
		fmt.Println("erastosthenes:")
		start := time.Now()
		sum := primes.Erastosthenes(treshold)
		fmt.Printf("Sum primes\t%v\n", sum)

		elapsed := time.Since(start)
		fmt.Printf("Primes took %s\n", elapsed)
	}
}
