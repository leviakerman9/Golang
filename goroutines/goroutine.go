package main

import (
	"fmt"
	"sync"
)

// goroutines is powerfull mechanism that spilt the task in multi light weight thread by which mutiple execute run paralelly
// Goroutines are a lightweight, concurrent execution unit in Go.
// They allow functions to run concurrently with the main program and other goroutines, making it efficient to handle multiple tasks simultaneously.
// Goroutines are managed by the Go runtime, making them significantly cheaper to create and manage than traditional threads

func print(i int, w *sync.WaitGroup){ // wait grp are passed in pointer format
	defer w.Done() //defer keyword means this func will excute at the end of func
	fmt.Println("task -> ",i)
}

func main(){

	// wait grp
	// there are three imp thing in wait grp add done wait
	/* add -> will add goroutine in wg
	   done -> remove goroutine fromo wg
	   wait -> wait till goroutine is completed	
	*/
	var wg sync.WaitGroup

	for i:=1; i<=10; i++{
		wg.Add(1)
		go print(i ,&wg)
	}


	// you can make inline function also
	// for i := 1; i <=10; i++ {
		
	// 	go func(i int){
	// 		fmt.Println(i)
	// 	}(i)
	// }

		// time.Sleep(time.Second *2) // wait time ye islie lagaya kuki jbtk hamara goroutines run hora hai tb tk
		// main function end ho jara h esliye wait lagaya hai main func ke end hone se pehle
	wg.Wait()
}

// this is the output
/*10
6
8
7
9
2
4
5
3
1*/ 
/*it shows that function print is divided in mutliple
 light weight thread which are executing parallel
 and completed in different manner they are not taking pause like simple function call execution which are one by one
*/