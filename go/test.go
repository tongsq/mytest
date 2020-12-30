package main

import "fmt"

//func modify(sls []int) {
//	sls[0] = 90
//}
//
//func main() {
//	var i int = 10
//	fmt.Println(interface{}(i).(int))
//
//	a := [3]int{89, 90, 91}
//	modify(a[:])
//	fmt.Println(a)
//
//}
var name string = "go"

func myfunc1() string {
	defer func() {
		name = "python"
	}()

	fmt.Printf("myfunc 函数里的name：%s\n", name)
	return name
}

func main() {
	myname := myfunc1()
	fmt.Printf("main 函数里的name: %s\n", name)
	fmt.Println("main 函数里的myname: ", myname)
}
