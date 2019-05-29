package main

import(
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	fmt.Println(now) 			// prints current date & time
	fmt.Println(now.Month())	// prints current month
	fmt.Println(now.Day()) 		// prints current number of day of the week
	fmt.Println(now.Weekday()) 	// prints current day

	oneDayLater := now.AddDate(0, 0, 1)
	oneMonthLater := now.AddDate(0, 1, 0)
	oneYearLater := now.AddDate(1, 0, 0)

	fmt.Println(oneDayLater) 	// prints next day with current time
	fmt.Println(oneMonthLater) 	// prints the next month with current time
	fmt.Println(oneYearLater) 	// prints the next year with current time

}

// https://golang.org/pkg/time/