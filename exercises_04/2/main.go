package main

import "fmt"

func main() {
	c := make(chan int)
	r := 10

	go func() {
		for i := 0; i < r; i++ {
			c <- i
		}
		//or
		// close(c)
	}()

	for i := 0; i < r; i++ {
		fmt.Println(<-c)
	}
	//or
	// for v := range c {
	// 	fmt.Println(v)
	// }
}
