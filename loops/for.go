package main

func main(){

	// while loop using for loop in go there is only for loop
	var i=0
	for i<3{
		println(i);
		i++
	}

	// for loop

	for i=0; i<5; i++{
		if i==3 {
			continue
		}
		println(i);
	}

	// using range keyword

	for i= range 11{

		
		println(i);
		if i==5{
			break
		}
	}
}