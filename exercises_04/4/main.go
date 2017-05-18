package main

import "fmt"

// func main() {
// 	c := factorial(4)
// 	for n := range c {
// 		fmt.Println(n)
// 	}
// }
//
// func factorial(n int) chan int {
// 	out := make(chan int)
// 	go func() {
// 		total := 1
// 		for i := n; i > 0; i-- {
// 			total *= i
// 		}
// 		out <- total
// 		close(out)
// 	}()
// 	return out
// }

func main() {
	c := gen(20)
	out := fact(c)
	for n := range out {
		fmt.Println(n)
	}
}

func gen(n int) chan int {
	out := make(chan int)
	go func() {
		for i := n; i > 0; i-- {
			out <- i
		}
		close(out)
	}()
	return out
}

func fact(in chan int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for n := range in {
			total *= n
		}
		out <- total
		close(out)
	}()
	return out
}
