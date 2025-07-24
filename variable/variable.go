package main

import "fmt"

// variable can be declared outside of main function also like.....

var topic string="training"
var topic1 ="golang"

// but shorthand syntax cannot be used outside

func main() {

	// first define then declare
	var name string
	var id int
	var isAdult bool
	var height float32
	

	name ="aditya"
	id=20
	isAdult=true
	height=5.92

	fmt.Println(name,id,isAdult,height)

	// define declare at same it and not using datatype it will know by itself it is string or any datatype

	var name2 ="aditya chauhan"
	var id2  = 2
	var isAdult2 = true
	var height2 = 10.444

	fmt.Println(name2,id2,isAdult2,height2)

	// shorthand syntax
	
	name3:="aditya singh chauhan"
	id3:= 3
	isAdult3:= true
	height3:=11.44

	fmt.Println(name3,id3,isAdult3,height3)

	fmt.Println("these variable are defined outside =",topic,topic1);

	
}