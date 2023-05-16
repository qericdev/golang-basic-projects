package main

import "fmt"

func main() {

	fmt.Println("Give a number to calculate the factorial")
	var n int
	fmt.Scanln(&n)

	fmt.Println("The factorial is:", factorial(n))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
