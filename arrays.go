package main

import(
	"fmt"
)

func main(){
	var colors [3] string
	colors[0]="Blue"
	colors[1]="Green"
	colors[2]="Yellow"

	fmt.Println(colors[0])
	fmt.Println(colors[1])
	fmt.Println(colors[2])
	fmt.Println(colors)

	var numbers = [3] int {1,2,3}

	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])
	fmt.Println(numbers)

	fmt.Println("Number of colors: ", len(colors))
	fmt.Println("Number of numbers: ",len(numbers))
}

