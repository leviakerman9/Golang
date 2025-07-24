package main

import "fmt"

type grade struct{
	marks int
}

type student struct {
	name   string
	class  string
	rollno int
	 grade // struct embeding in another struct (like grade struct is embeded in student struct)
}

func(s *student) update(no int){
	s.rollno=no
}
func(s *student) updatename(str string){
	s.name=str
}



func main() {

	g1:=grade{
		marks: 85,
	}

	g2:=grade{
		marks: 90,
	}

	var s1 = student{
		name:   "aditya",
		class:  "A",
		rollno: 1,
		grade:g1,
	}

	s2 := student{
		name:   "harsh",
		class:  "A",
		rollno: 2,
		grade: g2,
	}

	// fmt.Println("first object ",s1)
	// fmt.Println("sec object ",s2)
	// fmt.Println("s1 rollno ",s1.rollno)
	// fmt.Println("s2 class",s2.class)

	// s1.update(5)
	// fmt.Println("s1 rollno changed",s1.rollno)

	// fmt.Println("s1 previous name",s1.name)
	// s1.updatename("riya")
	// fmt.Println("s1 name changed",s1.name)
	
	// fmt.Println("s2 previous rollno ",s2.rollno)
	// s2.update(10)
	// fmt.Println("s1 rollno changed",s2.rollno)

	// fmt.Println("s2 previous name",s2.name)
	// s2.updatename("priya")
	// fmt.Println("s2 name changed",s2.name)
	// fmt.Println(s2.name)

	//  you can also make struct like this if you only want
	//  to use it once and not required to make multiple instance

	// language := struct {
	// name string
	// isGood bool
	// }{"golang",true}
	 
	//  inline struct 

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1.marks)
	fmt.Println(s2.marks)


}