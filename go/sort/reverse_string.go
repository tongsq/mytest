package main

import "fmt"

func ReverseString(b []byte, start, end int) []byte {
	for start < end {
		t := b[start]
		b[start] = b[end]
		b[end] = t
		start++
		end--
	}
	return b
}

func LeftRotateString(s string, m int) string {
	b := []byte(s)
	n := len(b)
	b = ReverseString(b, 0, m-1)
	b = ReverseString(b, m, n-1)
	b = ReverseString(b, 0, n-1)
	return string(b)
}

func main() {
	s := "abc123456"
	s = LeftRotateString(s, 4)
	fmt.Println(s)
	var i uint = 0
	j := (int)(int(i) >> 1)
	println(j)
}
