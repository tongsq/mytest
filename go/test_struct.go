package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  uint32
}

type Company struct {
	Name string
	addr string
}

type Staff struct {
	Name string
	Age  int
	Company
}

func (company Company) SayCompany() {
	fmt.Println("company is ", company.Name)
}

func (company Company) String() string {
	b, err := json.Marshal(company)
	fmt.Println(err)
	return string(b)
}

func (person *Person) Say() {
	fmt.Println(person.Name)
}

func (s *Staff) String() string {
	return s.Name
}

func (s *Staff) Modify() {
	s.Name = "bbbbbb"
}

func main() {
	var p Person = Person{"小A", 18}
	p.Say()
	myCom := Company{
		Name: "alibaba",
		addr: "hangzhou",
	}
	staffInfo := Staff{
		Name:    "小明",
		Age:     18,
		Company: myCom,
	}
	fmt.Println(staffInfo.Name, staffInfo.Company.Name)
	staffInfo.SayCompany()
	staffInfo.Company.SayCompany()
	fmt.Println(staffInfo)
	staffInfo.Modify()
	fmt.Println(staffInfo)
	staffInfo.addr = "123abc"
	fmt.Println(staffInfo.Company.addr)
}
