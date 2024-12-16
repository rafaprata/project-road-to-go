package main

import "fmt"

func main() {
	// var i int
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range ints {
		fmt.Println(v)
	}
}
