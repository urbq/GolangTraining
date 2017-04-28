package main

import "fmt"

func main() {
	highest := func(list ...int) int {
		var high int
		for _, n := range list {
			if n > high {
				high = n
			}
		}
		return high
	}
	fmt.Println(highest(5, 8, 4, 2, 7, 555, 7, 2))
}
