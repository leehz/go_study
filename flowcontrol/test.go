package main

import "fmt"

func main() {
	sum := 1
	ss := 1
	for sum < 1000 {
		sum += sum
	}

	for ss < 1000 {
		ss = ss * 2
	}
	fmt.Println(sum)
	fmt.Println(ss)
}
