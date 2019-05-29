package main

import(
	"fmt"
)

type User struct {
	Name string
	Age int
}

func main(){
	user := User{"Chan",32}
	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Printf("%+v",user)
}