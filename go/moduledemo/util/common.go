package util

//var num int
//Sum add args
func Sum(args ...int) (sum int) {
	for _, val := range args {
		sum += val
	}
	return
}
