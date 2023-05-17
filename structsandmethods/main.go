package main

import "fmt"

type rectangle struct {
	width  int
	height int
}

func main() {

	fmt.Println("Hi!, enter the width of the rectangle")
	var width int
	fmt.Scanln(&width)
	fmt.Println("Hi!, enter the height of the rectangle")
	var height int
	fmt.Scanln(&height)

	r := rectangle{width, height}

	fmt.Println("Area of rectangle is:", r.area())
	fmt.Println("Perimeter of rectangle is:", r.perim())

}

func (r rectangle) area() int {
	return r.width * r.height
}

func (r rectangle) perim() int {
	return 2*r.width + 2*r.height
}
