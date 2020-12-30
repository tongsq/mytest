package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int
}

func (p Person) SayBye() string {
	return p.name
}

func (p Person) SayHello() (string, string) {
	return "Hello", "world"
}

func (p Person) Say(word string) (string, string) {
	return word, "ok"
}

type SayInterface interface {
	Say(string2 string) (string, string)
}

func main() {
	p := &Person{"aaaa", 27}

	t := reflect.TypeOf(p)
	//dst := (*SayInterface)(nil)
	var dst *SayInterface
	dstType := reflect.TypeOf(dst).Elem()
	if t.Implements(dstType) {
		fmt.Println("p implements SayInterface")
	}
	v := reflect.ValueOf(p)

	for i := 1; i < v.NumMethod(); i++ {
		fmt.Printf("调用第 %d 个方法：%v ，调用结果：%v\n",
			i+1,
			t.Method(i).Name,
			v.Method(i).Call(nil))
	}
	fmt.Println(v.MethodByName("SayHello").Call(nil))
	name := reflect.ValueOf("word test")
	input := [1]reflect.Value{name}
	//var input []reflect.Value
	//input = append(input, name)
	fmt.Println(v.MethodByName("Say").Call(input[:]))
}
