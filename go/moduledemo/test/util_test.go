package test

import "testing"
import "moduledemo/mypackage"

func TestAdd(t *testing.T) {
	sum := mypackage.Add(1, 2)
	if sum == 3 {
		t.Log("the add result is ok")
	} else {
		t.Fatal("the add result is wrong")
	}
}
