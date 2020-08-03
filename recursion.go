package main

import "fmt"

func main()  {
	var i uint64 = 15
	fmt.Println(Factorial(i))

	var n int
	for n = 0; n < 10; n++ {
		fmt.Println("%d\t", fibonacci(n))
	}
}

// 一个数的阶乘
func Factorial(n uint64) (result uint64)  {
	if (n > 0) {
		result = n * Factorial(n - 1)
		return result
	}

	return 1
}

// 斐波那契数列
func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n - 2) + fibonacci(n - 1)
}
