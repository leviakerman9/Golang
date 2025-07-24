package main

import (
	"fmt"
)

// constant can also be defined outside of function

func main() {

	// constants value cannot be changed
	const n=5

	// constants can be grouped

	const(
		name ="aditya"
		 id=1
	)
	fmt.Println(n)
	fmt.Println(name,id);
}