package main

import "fmt"

func main() {

	fmt.Println("Insert the number of elements to sort")
	var n int
	fmt.Scanln(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Insert the %d element: \n", i+1)
		fmt.Scanln(&arr[i])
	}

	fmt.Println("Choose the sorting method: 1. Bubble 2. Selection")
	var method int
	fmt.Scanln(&method)

	switch method {
	case 1:
		fmt.Println("Bubble sort: ", sortWithBubble(arr))
	case 2:
		fmt.Println("Selection sort: ", sortWithSelection(arr))
	default:
		break
	}
}

func sortWithBubble(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func sortWithSelection(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		temp := arr[minIndex]
		arr[minIndex] = arr[i]
		arr[i] = temp
	}
	return arr
}
