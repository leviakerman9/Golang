package main

import "fmt"

// Go (Golang) does not have a built-in enum keyword like C/C++, Java, or Rust. 
// However, you can simulate enums using constants and iota.


type OrderStatus string

const (
	Received  OrderStatus = "RECEIVED"
	Confirmed OrderStatus = "CONFIRMED"
	Prepared  OrderStatus = "PREPARED"
	Delivered OrderStatus = "DELIVERED"
)

func status(str OrderStatus) {
	fmt.Println("The order has been -> ",str)
}

func main() {

	status(Confirmed)

}