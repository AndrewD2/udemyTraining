package main

import "fmt"

func largest(x ...int) {

	var largest int

	for _, v := range x {
		if v > largest {
			largest = v
		}
	}

	fmt.Println("The largest value is: ", largest)
}

func main() {


	largest(82, 100, 5000, 1, 5)

}
