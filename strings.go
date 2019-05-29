package main

import(
	"fmt"
	"strings"
)

func main(){

	stra := "hello" // Implicitly  Type
	var strb string = "from" // Static Type
	strc := "Strings"

	fmt.Println(stra,strb,strc)
	fmt.Printf("%v,%v,%v",stra,strb,strc)	// outputing string with Verbes
	fmt.Printf("%T\n",stra)	// outputing string value Types

	fmt.Println(strings.ToUpper(stra)) // outputs strings to Uppercase
	fmt.Println(strings.Title(stra)) // outputs Title Case / First letter to uppercase

	const a string = "HELLO!"
	var b string = "hello!"

	fmt.Println("Is Equal = ", (a == b )) // returns False 
	fmt.Println("Is Equal = ", strings.EqualFold(a,b)) // returns True

	fmt.Println(strings.Count("hello", ""))
}


// know more about strings : https://golang.org/pkg/strings/