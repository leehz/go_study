package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	ss := 0
	for j := 0; j <= 100; j++ {
		ss += j
	}
	fmt.Println(sum)
	fmt.Println(ss)
}
