package main

import "fmt"

func main() {

	// by defalut contain zeros
	var nums = [4]int{1, 3, 4, 5}

	// shorthand syntax
	// nums:=[4]int{1,2,3,4}
	fmt.Println(nums)

	//  contains spaces by default
	var val =[3]string{"aditya","singh","chauhan"}

	fmt.Println(val)


	// bydefalut array contains false
	var val2=[2]bool{true,false}

	fmt.Println(val2)

	// 2d array
	var arr=[2][2]int {{1,1},{2,2}}
	fmt.Println(arr)

	fmt.Println(len(arr))
}

// fixed size use array means you know size ,that is predictable
// memory optimization
// constant time access
// array are fixed size not dynamic