package main

import "fmt"

func main() {
	fmt.Println("Insert the string that needs to be reversed:")
	var input string
	fmt.Scanln(&input)

	fmt.Println("The reversed string is:", reversed(input))

}

func reversed(input string) string {
	var reversed string
	for i := len(input) - 1; i >= 0; i-- {
		reversed += string(input[i])
	}
	return reversed
}
