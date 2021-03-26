package test

import "fmt"
import "strconv"
import "testing"

var num = 8568452
var numstr = "8568452"

//测试 strconv.ParseInt
func BenchmarkStrconvParseInt(b *testing.B) {
	num64 := int64(num)
	for i := 0; i < b.N; i++ {
		x, err := strconv.ParseInt(numstr, 10, 64)
		if x != num64 || err != nil {
			b.Error(err)
		}
	}
}

//测试 strconv.Atoi
func BenchmarkAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x, err := strconv.Atoi(numstr)
		if x != num || err != nil {
			b.Error(err)
		}
	}
}

//测试 fmt.Sscan
func BenchmarkFmtSscan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x int
		n, err := fmt.Sscanf(numstr, "%d", &x)
		if n != 1 || x != num || err != nil {
			b.Error(err)
		}
	}
}
