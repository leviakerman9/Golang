package main

import "fmt"

// generics can be used to pass any type or favaourable typr in function which can be restricted if we specify type

func print[T any](slice []T){   //you can here specify favourable type also like (T int | string ....)

	for _,key:= range slice{
		fmt.Println(key)
	}
}
// can also be used with struct
func main(){

	arr1:=[]int{1,2,3,4}
	arr2:=[]string{"go","lang","type","script"}

	print(arr1)
	print(arr2)

}