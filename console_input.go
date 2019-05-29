package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main(){
	// var s string             // Declearing a empty variable as string
	// fmt.Scanln(&s)           // Scanln(&s) scans the value of console input and stores tbe value in the empty variables
	// 						    // it also scans single input 
	// fmt.Println("hello", s)  // prints the value of the variable
	// fmt.Printf("%T",s)       // prints the type of variables

	reader :=  bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your Name: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(" Hello " +  str)
	
	fmt.Print("Enter Your Number: ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(f)
	}
	
}
