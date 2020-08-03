package main

import "fmt"

var intval = 456

const (
	b string = "abc"
	c int = 1
	d = 2
	e = iota
	f
	g
)

func main() {
    var intval int
    intval, intval1 := 123, 456
	fmt.Println(intval, intval1)
	const c = 2
	fmt.Println(c)
	fmt.Println(e, f, g)
}
