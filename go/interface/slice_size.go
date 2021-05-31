package main

import "fmt"

func main() {
	s1 := []int{}
	fmt.Println(s1, len(s1), cap(s1), s1)
	s1 = append(s1, 1)
	fmt.Println(&s1, len(s1), cap(s1), s1)
	s1 = append(s1, 2)
	fmt.Println(&s1, len(s1), cap(s1), s1)
	s1 = append(s1, 3)
	fmt.Println(&s1, len(s1), cap(s1), s1)
	s1 = append(s1, 4)
	fmt.Println(&s1, len(s1), cap(s1), s1)
}
