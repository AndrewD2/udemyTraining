package main

import "fmt"

func half(v int) (float64, bool) {

	half := float64(v) / 2
	even := false
	if v%2 == 0 {
		even = true
	}

	return half, even

}

func main() {

	x, y, z := 1, 2, 3

	fmt.Println(half(x))
	fmt.Println(half(y))
	fmt.Println(half(z))
}
