package main

import "fmt"

func main() {

	var m=map[int]string{}

	m[0]="aditya"
	m[1]="chauhan"

	var arr = []int{}

	arr = append(arr, 1, 2, 3, 4, 5)

	// iterate using for loop

	for i := 0; i < len(arr); i++ {

		fmt.Println(arr[i]);
	}

	for index,k:= range arr{
		fmt.Println(k,index)
	}

	for k,v:= range m{
		fmt.Println(k,v);
	}


}