package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Animal interface {
	Grow()
	Move(string) string
	Say()
}

type Cat struct {
	Name    string
	Age     uint32
	Address string
}

func (cat *Cat) Grow() {
	cat.Age = cat.Age + 1
}

func (cat *Cat) Move(new string) (old string) {
	old = (*cat).Address
	cat.Address = new
	return
}
func (cat *Cat) Say() {
	fmt.Println("cat say")
}

type Dog struct{}

func (dog *Dog) Grow() {}

func (dog *Dog) Move(new string) (old string) {
	old = "old"
	return
}
func (dog *Dog) Say() {
	fmt.Println("dog say")
}

func main() {
	myCat := Cat{"Little C", 2, "In the house"}
	animal, ok := interface{}(&myCat).(Animal)
	fmt.Println(reflect.TypeOf(animal))
	fmt.Printf("%v, %v\n", ok, animal)
	animal.Grow()
	fmt.Printf("%v\n", animal)
	fmt.Println(myCat.Move("new"))
	dog := Dog{}
	var animals []Animal
	animals = append(animals, &myCat)
	animals = append(animals, &dog)
	for _, animal := range animals {
		animal.Say()
	}

	a := 10
	b := "hello"
	c := true
	myfunc(a, b, c)

	var str1 string = strconv.Itoa(a)
	// var i int = int(str1)
	fmt.Printf("type: %T, value: %s\n", str1, str1)
	t, ok := interface{}(str1).(string)
	fmt.Println(t)
}

func myfunc(ifaces ...interface{}) {
	for _, iface := range ifaces {
		fmt.Println(iface)
	}
}
