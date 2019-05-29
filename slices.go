// An array has a fixed size. A slice,
// on the other hand, is a dynamically-sized, 
// flexible view into the elements of an array.
// In practice,slices are much more common than arrays

package main

import(
	"fmt"
)


func main(){
	var str = [] string {"Chan", "Bony"}
	fmt.Println(str)

	var num = [] int {}
	num = append(num,1)
	num = append(num,2)
	num = append(num,3)

	fmt.Println(num)    			 //	outputs [1,2,3]
	fmt.Println(num[0])				 //	outputs 1
	fmt.Println(len(num))			 //	outputs 3
	fmt.Println(num[:len(num)-1])    //	outputs [1,2] , removes the last value of the array

	var inte = make([] int, 2)
	inte[0]= 2
	inte[1]= 2

	fmt.Println(inte)

}