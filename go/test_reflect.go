package main

import (
	"fmt"
	"reflect"
)

type Person1 struct {
	Name string
	age  int
}

func (p Person1) SayName() {
	fmt.Println("my name is ", p.Name)
}
func (p *Person1) SayAge() {
	fmt.Println("my age is ", p.age)
}

func main() {
	p := Person1{"aa", 18}
	v := reflect.ValueOf(&p)
	v.MethodByName("SayName").Call(nil)
	for i := 0; i < v.NumMethod(); i++ {
		v.Method(i).Call(nil)
	}
	var name string = "test"
	v1 := reflect.ValueOf(&name)
	v2 := v1.Elem()
	v2.SetString("test2")
	fmt.Println(name)
	v3 := reflect.ValueOf(&name).Elem()
	fmt.Println(v3.CanSet())
}
