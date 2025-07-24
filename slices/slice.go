package main

import (
	"fmt"
	"slices"
)

// slices are dynamix array like vector in c++
// most used construct
// it contains useful methods

func main(){

	// unintialized slice in default nil

	// var nums[]int

	// var num = make([]int,0,5)
	// use make function when element are not known so you can add later
	// 2 varible is length of slice it creates slices with 0 like if length is 2 then slice is[0,0,1,.....] 
	// first 2 no will be zero
	// 3 variable is capacity
	// num = append(num, 1)
	// num = append(num, 2)
	// num = append(num, 3)
	// num = append(num, 4)

	// fmt.Println(num);
	// fmt.Println(len(num))
	// fmt.Println(num[2])

	// copy function to copy one slice into other

	var num1=make([]int,0,5)
	num1 = append(num1, 2,3,4,6)
	var num2=make([]int,len(num1))

	copy(num2,num1)

	fmt.Println(num2)

	var num3=[]int{1,2,3,4}
	
	// slice operator
    fmt.Println(num3[0:2]) 	// in slice it will return value from index 0 to index 1 it retrun value one less from last index

	fmt.Println(slices.Equal(num1,num2))

	// 2d slices
	var nums=[][]int{{1,2,3},{4,5,6}}
	fmt.Println(nums)

	var num5=[]int{}

	num5 = append(num5, 2,3)
	fmt.Println(num5)
}