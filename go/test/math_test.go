package test

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(1, 2, 3)
	if sum == 6 {
		t.Log("TestAdd success")
	} else {
		t.Fatal("TestAdd failed")
	}
}

func Add(args ...int) (sum int) {
	for _, v := range args {
		sum = sum + v
	}
	return
}
