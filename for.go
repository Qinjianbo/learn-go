package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	// 每次都要拼接新字符串sep，然后赋值给s
	// 如果量大的话，代价会很大
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	// 使用strings包的join方法会更高效
	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println(os.Args[1:])

	fmt.Println(os.Args[0])

	seq := " "
	for index, arg := range os.Args[1:] {
		index_str := strconv.Itoa(index)
		fmt.Println(index_str + seq + arg)
	}
}
