package main

import "fmt"

func main() {
	v := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for n := range v {
		if n%2 == 0 {
			fmt.Printf("%v is even\n", n)
		} else {
			fmt.Printf("%v is odd\n", n)
		}
	}
}
