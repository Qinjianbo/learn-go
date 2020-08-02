package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5, 6, 7}
	PrintSliceContent(slice)
	fmt.Println(Sum(slice))

	slice1 := map[string]string{"name": "胡巴", "age": "28"}
	fmt.Println(slice1)
	for k, v := range slice1 {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range 用来枚举Unicode 字符串, 第一个参数是字符串的索引， 第二个是字符（Unicode的值）本身
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

// 思考：这里怎么能接收所有类型的slice
func PrintSliceContent(s []int) {
	for index, value := range s {
		fmt.Printf("index:%d => value:%d\n", index, value)
	}
}

func Sum(s []int) int {
	sum := 0
	for _, value := range s {
		sum += value
	}

	return sum
}
