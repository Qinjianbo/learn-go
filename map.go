package main

import "fmt"

func main() {
	// 声明变量 默认 map 是nil
	var map1 map[string]int
	fmt.Println(map1 == nil) // true
	fmt.Println(map1) // map[]

	map1 = map[string]int{}
	fmt.Println(map1 == nil) // false
	fmt.Println(map1) // map[]

	var map2 = make(map[string]int)
	fmt.Println(map2 == nil) // false
	fmt.Println(map2) // map[]

	map3 := make(map[string]int)
	fmt.Println(map3 == nil) // false
	fmt.Println(map3) // map[]

	countryCapitalMap := map[string]string{
		"France": "巴黎",
		"Italy" : "罗马",
		"Japan" : "东京",
		"India" : "新德里" }
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	// 查看元素在集合中是否存在
	capital, ok := countryCapitalMap["American"]
	fmt.Println(capital)
	fmt.Println(ok)
	if (ok) {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	// 使用delete 删除元素
	delete(countryCapitalMap, "Japan")
	for country := range countryCapitalMap {
		fmt.Println(country, "的首都是", countryCapitalMap[country])
	}
}
