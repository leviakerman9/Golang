package main

import "fmt"

func change(num *int){

	*num=6

}

func main() {

	var num = 5

	var ptr = &num //stroing num add
	fmt.Println(*ptr) //return value at num add dereference

	*ptr = 6

	fmt.Println(*ptr)

	*ptr=*ptr+1

	fmt.Println(*ptr)  // pointer stores or point to value present at num address
	fmt.Println(ptr)   //ptr stores num address
	fmt.Println(&ptr)  //pointer address

	var n=5

	change(&n)  //pass by reference
	fmt.Println(n)

}