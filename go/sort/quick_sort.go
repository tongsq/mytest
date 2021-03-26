package main

import "fmt"

func partition(list []int, low, high int) int {
	pivot := list[low]
	for low < high {
		for low < high && pivot <= list[high] {
			high--
		}
		list[low] = list[high]
		for low < high && pivot >= list[low] {
			low++
		}
		list[high] = list[low]
	}
	list[low] = pivot
	return low
}

func QuickSort(list []int, low, high int) {
	if low < high {
		//位置划分
		pivot := partition(list, low, high)
		//排序左边部份
		QuickSort(list, low, pivot-1)
		//排序右边部份
		QuickSort(list, pivot+1, high)
	}
}

func main() {
	list := []int{2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54, 9, 100, 30, 34}
	QuickSort(list, 0, len(list)-1)
	fmt.Println(list)
}
