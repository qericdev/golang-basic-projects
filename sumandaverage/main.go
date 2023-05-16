package main

import "fmt"

func main() {

	fmt.Println("Enter the length of the array: ")
	var n int
	var sum int
	fmt.Scanln(&n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Println("Enter the number at index ", i, ": ")
		fmt.Scanln(&arr[i])
		sum += arr[i]
	}

	average := float64(sum) / float64(n)

	fmt.Println("Sum of the array is: ", sum)
	fmt.Println("Average of the array is: ", average)
}
