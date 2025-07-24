package main

import (
	"fmt"
	"sync"
)

type value struct {
	val int
	mu  sync.Mutex
}

func (v *value)inc(w *sync.WaitGroup) {
	defer func(){
		v.mu.Unlock()
		w.Done()
	}()
	
	v.mu.Lock()
	v.val++

}

func main() {

	var wg sync.WaitGroup

	v1:=value{val:0}

	for i:=0; i<1000; i++{
	wg.Add(1)
	

	go v1.inc(&wg)

	}

	wg.Wait()
	fmt.Println(v1.val)

}