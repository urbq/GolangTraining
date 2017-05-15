package primes

import (
	"fmt"
	"math"
)

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
	return GetPrimesForSlice(numbers...)
}

func GetPrimesForSlice(numbers ...int) []int {
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

func Test(treshold int) []int {
	//berechne quadratqurzel aus grenze
	sqrt := math.Sqrt(float64(treshold))

	//1.erstelle eine Liste mit allen Zahlen von 2 bis N
	//...am besten nur ungerade
	var numbers []int
	numbers = append(numbers, 2)
	t := 3
	for t <= treshold {
		numbers = append(numbers, t)
		t += 2
	}

	fmt.Println(sqrt)
	// fmt.Print("start")
	// fmt.Println(numbers)

	for i := 0; i < len(numbers); i++ {
		// fmt.Println("i[", i, "]=", numbers[i])
		//2.finde die nächste Zahl p, die nicht ausgeschlossen wurde. Dies ist eine Primzahl. Ist sie größer als die Wurzel aus N, gehe zu 5
		if isPrime(numbers[i]) {
			//3.Markiere alle Vielfachen von p als “keine Primzahl”
			if float64(numbers[i]) < sqrt {
				continue
			}
			for j := i; j < len(numbers); j++ {
				// fmt.Println("\tj[", j, "]=", numbers[j])
				if numbers[j] > numbers[i] && numbers[j]%numbers[i] == 0 {
					numbers = append(numbers[:j], numbers[j+1:]...)
					// fmt.Println(numbers)
				}
			}
		} else {
			numbers = append(numbers[:i], numbers[i+1:]...)
			i--
			// fmt.Println(numbers)
		}
	}
	// fmt.Print("end")
	// fmt.Println(numbers)
	return numbers
}

func Test2(treshold int) []int {
	primVec := []int{2}

	prim := 3
	for prim < treshold {
		wurzel := int(math.Sqrt(float64(prim)))
		fmt.Println(wurzel)
		istPrim := 1
		teiler := 3
		for teiler < wurzel {
			if prim%teiler == 0 {
				istPrim = 0
			}
			teiler += 2
		}
		if istPrim == 1 {
			primVec = append(primVec, prim)
		}
		prim += 2
	}
	return primVec
}
