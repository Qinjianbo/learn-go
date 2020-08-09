package main

import "fmt"

func main() {
	var i int
	var doub float64
	var str string
	var comp complex64
	var a []int
	var p *int
	var m map[string] int
	var f func(string) int
	var e error
	fmt.Printf("%v %v %q %v %q %v %v %v %v", i, doub, str, comp, a, p, m, f, e)
}
