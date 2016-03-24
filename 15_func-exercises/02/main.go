package main

import "fmt"

func main() {

	func(v int) func(float64, bool) {

		half := float64(v) / 2
		even := false
		if v%2 == 0 {
			even = true
		}

		return half, even

	}
	x, y, z := 1, 2, 3

	fmt.Println((x))
	fmt.Println((y))
	fmt.Println((z))
}
