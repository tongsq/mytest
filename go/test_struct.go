package main

import "fmt"

type Person struct {
	Name string
	Age  uint32
}

type Company struct {
	name string
	addr string
}

type Staff struct {
	name string
	age  int
	Company
}

func (company Company) SayCompany() {
	fmt.Println("company is ", company.name)
}

func (person *Person) Say() {
	fmt.Println(person.Name)
}
func main() {
	var p Person = Person{"小A", 18}
	p.Say()
	myCom := Company{
		name: "alibaba",
		addr: "hangzhou",
	}
	staffInfo := Staff{
		name:    "小明",
		age:     18,
		Company: myCom,
	}
	fmt.Println(staffInfo.name, staffInfo.Company.name)
	staffInfo.SayCompany()
}
