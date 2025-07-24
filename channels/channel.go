package main

import (
	"fmt"
	"sync"
	"time"
)

func sender(ch1 chan string,w *sync.WaitGroup){

	defer w.Done()
	msg:="hii"

	for i:=0; i<10; i++{
		ch1 <-msg
		time.Sleep(time.Second *2)
	}

	close(ch1)
}

func forwarder(ch2 ,ch1 chan string,w *sync.WaitGroup){
	defer w.Done()

	for val:= range ch1{
		ch2 <- val
	}

	close(ch2)
}

func receiver(ch2 chan string,w *sync.WaitGroup){
	defer w.Done()

	for val:= range ch2{
		fmt.Println("received",val)
	}
}

func main() {

	ch1:=make(chan string)
	ch2:=make(chan string)

	var wg sync.WaitGroup

	wg.Add(3)

	go sender(ch1,&wg)
	go forwarder(ch2,ch1,&wg)
	go receiver(ch2,&wg)

	wg.Wait()
	fmt.Println("all msgs are printed")



	// how to create channel
/*	messageChan := make(chan string)

	messageChan <- "hi" //sending end fo pipe | channel are blocking

	msg:= <- messageChan //receiving end

	fmt.Println(msg)

	but this will error of deadlock because channel are blocking
	it will not send msg in channel pipe till there any to receive msg

	*/
}