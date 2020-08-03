package main

import "fmt"

type Pet interface {
	Name() string
	Age() uint8
}

type Dog struct {
	name string
	age  uint8
}

func (dog Dog) Name() string {
	return dog.name
}
func (dog Dog) Age() uint8 {
	return dog.age
}

func main() {
	myDog := Dog{"Little D", 3}
	_, ok1 := interface{}(&myDog).(Pet)
	_, ok2 := interface{}(myDog).(Pet)
	fmt.Printf("%v, %v\n", ok1, ok2)
	var str1 string = "hello"
	p1 := &str1
	var arr1 []int
	fmt.Printf("str:%s, p1: %X, p: %p %x, arr1:%p", *p1, p1, &str1, &str1, arr1)
}
