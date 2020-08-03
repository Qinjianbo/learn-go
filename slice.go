package main

import "fmt"

func main() {
	var slice1 []int;
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Printf("%p\n", slice1)

	var slice2 []int = make([]int, 10)
	fmt.Println(slice2)
	fmt.Println(len(slice2))

	slice3 := make([]int, 20)
	fmt.Println(slice3)

	slice4 := make([]string, 20)
	fmt.Println(slice4)
	fmt.Printf("%#v\n", slice4)

	slice5 := []int {12, 13, 14}
	fmt.Printf("%#v\n", slice5)
	fmt.Println(slice5[:])
	fmt.Println(slice5[:2])
	fmt.Println(slice5[1:])
	fmt.Println(slice5[0:1])
	fmt.Printf("len:%d,cap:%d,slice:%v\n", len(slice5), cap(slice5), slice5)

	slice6 := slice5[:2]
	fmt.Printf("len:%d,cap:%d,slice:%v\n", len(slice6), cap(slice6), slice6)
}
//各行输出结果
//[]
//0
//0x0
//[0 0 0 0 0 0 0 0 0 0]
//10
//[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
//[                   ]
//[]string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
//[]int{12, 13, 14}
//[12 13 14]
//[12 13]
//[13 14]
//[12]
//len:3,cap:3,slice:[12 13 14]
//len:2,cap:3,slice:[12 13]
