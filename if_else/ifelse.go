package main

import "fmt"

func main() {
	// if else 
	var age int
	fmt.Scan(&age)

	if age >18 {
		println("adult")
	}else if age>=12 && age<18{
		println("teenager")
	}else{
		println("teen")
	}

	// go doesnot have ternary operator

	

}