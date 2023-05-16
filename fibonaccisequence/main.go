package main

import "fmt"

func main() {

	fmt.Println("Gives a number n to generate the first n terms of the fibonacci series")
	var n int
	fmt.Scanln(&n)
	counter := 0
	val := 0

	var fibonacci []int = make([]int, n)

	for counter < n {
		fibonacci[counter] = val

		if counter > 0 {
			val = fibonacci[counter] + fibonacci[counter-1]
		} else {
			val++
		}
		counter++
	}

	fmt.Println("The serie is:", fibonacci)
}
