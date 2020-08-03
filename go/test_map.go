package main

import "fmt"

func main() {
	mm := map[int]string{1: "a", 2: "b", 3: "c"}
	b := mm[2]
	fmt.Println(b)
	mm[2] = b + "b"
	fmt.Println(b)
	fmt.Println(mm[2], mm)
	delete(mm, 2)
	fmt.Println(mm)
	val, ok := mm[3]
	if ok {
		fmt.Println("get index 2 ok,", val)
	} else {
		fmt.Println("get index 2 error")
	}
	var map1 map[string]string = make(map[string]string)
	map1["key1"] = "value2"
	fmt.Println(map1, map1["key1"])
	var arr1 []interface{} = make([]interface{}, 0, 10)
	arr1 = append(arr1, "abc")
	arr1 = append(arr1, 123)
	fmt.Println(arr1)
	fmt.Printf("%p \n", arr1)
	value1 := arr1[0]
	fmt.Println(value1)
	func1(map1)
	fmt.Printf("%v\n", map1)
	arr2 := [...]int{1, 2, 3}
	func2(&arr2)
	fmt.Println(arr2)
	s1 := make([]int, 5, 10)
	func3(s1)
	fmt.Println(s1)
}

func func1(mm map[string]string) {
	mm["abc"] = "abcd"
}

func func2(arr *[3]int) {
	arr[0] = 456
}

func func3(s1 []int) {
	s1[0] = 123
	fmt.Println(cap(s1), len(s1))
	fmt.Printf("%p\n", &s1)
	s1 = append(s1, 123)
	fmt.Println(s1, cap(s1), len(s1))
	fmt.Printf("%p\n", &s1)
}
