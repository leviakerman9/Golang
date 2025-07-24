package main

import (
	"fmt"
	"time"
)

func main() {

	// switch constructor(simple switch)

	var i,t int

	fmt.Scan(&i)
	fmt.Scan(&t)


	switch i{

	case 1: 
	println("case 1")
	
	case 2: 
	print("case 2")
	
	case 3: 
	println("case 3")

	default:
		println("other")
	}

	// multiple condition switch

	switch time.Now().Weekday(){

	case time.Saturday,time.Sunday:println("holiday")

	default:println("this is school day")
	}

}