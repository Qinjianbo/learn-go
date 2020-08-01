package main

import "fmt"

type Student struct{
	name string
	age  int
	sex  int
	no   string
}
func main() {
	var st Student
	st.name = "huba"
	st.age  = 28
	st.sex  = 1
	st.no   = "20120709"
	fmt.Printf("%#v\n", st)
	fmt.Printf("%v\n", st)
	fmt.Printf("%T\n", st)
	fmt.Printf("%%\n")
	fmt.Printf("%t\n", st)
	fmt.Println(st)
	fmt.Println(st.name)

	st1 := Student{name: "huba"}
	fmt.Println(st1)

	st2 := Student{"huba", 28, 1, "20120901"}
	fmt.Println(st2)
}
