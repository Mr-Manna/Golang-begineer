package main

import(
	"fmt"
)

func main(){
	a := "base"
	b := "ball"

	x := 5
	y := 6

	fmt.Println(a + b)  // Adds two variables
	fmt.Println(a,b)	// Only prints variables
	fmt.Println(x + y)  // Adds two variables


	length := len(a) // Length of string
	fmt.Println(length)

	fmt.Printf("value of a is : %v %v",a,b) // Printf function allows to add placeholders in a string for variables using %v
	//fmt.Printf("vaue of x as float number is : %.2f",float64(x))//  %f to insert float value
	//fmt.Printf("vaue of x as float number is : %T", x )//  %T to insert type of  value
}