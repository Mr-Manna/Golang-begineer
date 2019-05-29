//  Pointers : points to another variable address

package main

import(
	"fmt"
)

func main(){
	var a int
	var b = &a
	
	if b != nil{
		fmt.Println(*b)
	}else{
		fmt.Println("Value is nil")
	}

	var p int = 45
	var q = &p

	if q != nil{
		fmt.Println(*q)
	}else{
		fmt.Println("Value is nil")
	}
}