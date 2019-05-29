// Maps are unordered list of key : value pair

package main

import (
	"fmt"
	"sort"
)

func main(){

	name := make(map[string]string)

	name["a"] = "Amit"
	name["b"] = "Bony"
	name["c"] = "Chan"
	name["d"] = "Damn"
	name["e"] = "Eon"

	fmt.Println(name)

	for _, value := range name {
		fmt.Printf("Hello, %v\n", value)
	}

	users := make([]string, len(name))

	i := 0
	for u := range name {
		users[i] = u
		i++
	}
	sort.Strings(users)
	fmt.Println("\nSorted : ")
	for val := range users {
		fmt.Println(name[users[val]])
	}
}
