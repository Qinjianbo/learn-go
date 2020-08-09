package main

import "fmt"

func main() {
	var age [10] int
	age[0] = 1
	age[2] = 2
	var sex [2] string
	sex[0] = "男"
	sex[1] = "女"

	var workFor = [...]string{"萌推","Dada","波奇"}

	fmt.Println(age, sex, workFor)
}
