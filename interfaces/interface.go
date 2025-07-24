package main

import "fmt"

// create interface

type paymenter interface{

	// declare methods here which should be defined by functions

	pay(amount int)
	paycount(name string)

}

type payment struct{
	gateway paymenter
}

func(p payment) makepayment(amount int){
	p.gateway.pay(amount)
}

func(p payment) paymentCount(name string ){
	p.gateway.paycount(name)
}
type stripe struct{
	count int
}

func(s stripe)pay(amount int){
	fmt.Println("making payment using stripe",amount)
}
func(s *stripe)paycount(name string){
	s.count++

	fmt.Println("payment done by -> ",name,",payments using stripe -> ",s.count)
}

type razorpay struct{}

func(r razorpay)pay(amount int){
	fmt.Println("making payment using razorpay",amount)
}





func main(){


	s1:=&stripe{
		count: 0,
	}
	// r1:=razorpay{}
	
	newPayment:=payment{
		gateway: s1,
		// gateway: r1,
	}

	newPayment.paymentCount("aditya")
		newPayment.makepayment(100)

	newPayment.paymentCount("harsh")
		newPayment.makepayment(100)

	newPayment.paymentCount("mithun")
		newPayment.makepayment(100)

	
	

}