package util

import "testing"

func TestSum(t *testing.T) {
	sum := Sum(1, 1, 1)
	if sum == 3 {
		t.Log("sum func is ok")
	} else {
		t.Fatal("sum func is wrong")
	}
	sum = Sum(1, 2, 3)
	if sum == 7 {
		t.Log("sum func is ok")
	} else {
		t.Fatal("sum func is wrong")
	}
}
