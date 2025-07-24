package main

import "fmt"

func main() {

	// create map using make function
	var m = make(map[int]string)

	m[1] = "aditya"
	m[2] = "singh"
	m[3] = "chauhan"

	fmt.Println(m)

	// var m=map[int]string{1:"a",2:"b"}

	value,ok:=m[1]

	if ok{
		fmt.Println("true",value)
	}else{
		fmt.Println("false")
	}

	// clear() will clear everything from map
	// delete(m,1) will deleted key pair
	// maps.Equal() function return bool value maps are equal or not
	 

}