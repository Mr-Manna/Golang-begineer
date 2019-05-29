//  	--------------  Notes  -----------------  	//
//  	Math Requires Same Types				 	//
//   	Example : Can't Add a INTEGR to a FLOAT     //
//  	Have to Convert the Types To Same Type	 	//
//  	Pkg: https://golang.org/pkg/math/   	 	//
//  	---------------       ------------------  	//

package main

import(
	"fmt"
	
)

func main(){

	x,y,z := 4,5,6 // implicitly typed integer values
	sum := x + y + z
	fmt.Printf( "Sum of 2 integer values : %v\n" , sum )

	var a float64 = 20.56 // static typed float values	
	var b float64 = 10.35 // static typed float values

	sum2 := a + b 
	fmt.Printf( "Sum of 2 floating values : %.2f\n" , sum2 )
	 
}