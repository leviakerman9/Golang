package main

import "fmt"

func add(a, b int) (int) {
	return a+b
}
func value() (string,int,float64) {
	return "aditya",100,5.6
}
// naked return(named return) -> where you give name to return value and simply return(only return)

func nakedreturn(x,y int)(sum,multiply int){ // here you are giving name to your return type 
	sum=x+y
	multiply=x*y
	return // this is naked return 
}
// and non-naked return -> where you simply return the value and do not give it a name

func non_naked(x,y int)(int,int){ //here you are only defining return datatype
	var sum=x+y
	var multiply=x*y
	return sum,multiply
}

// variadic function which can take x no of inputs

func plus(x ...int) int{

	total:=0

	for _,num:= range x{

		total=total+num
	}

	return total

}

func main() {

	// in go u can return multiple values

	ans:=add(2,3)
	fmt.Println(ans)
	fmt.Println(value())
	fmt.Println(nakedreturn(3,4))
	fmt.Println(non_naked(3,4))

	var arr=[]int {1,2,3,4,5}
	fmt.Println(plus(2,2,2,2,2))
	fmt.Println(plus(arr...))

}


// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

// In this example, we shortened

// x int, y int

// to

// x, y int