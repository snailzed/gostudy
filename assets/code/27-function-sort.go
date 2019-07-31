package main

import "fmt"

func main() {
	var arr []int = []int{1, 4, 35, 7, 93, 12, 0, 4, 6, 721, 2, 9}
	fmt.Println(BubbleSort(arr))
	fmt.Println(SelectSort(arr))
	fmt.Println(InsertSort(arr))
	fmt.Println(QuickSort(arr))
}

//冒泡排序算法,相邻比较
func BubbleSort(arr []int) (res []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] //互换
			}
		}
	}
	return arr
}

//选择排序
func SelectSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		k := i
		tmp := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if tmp > arr[j] {
				k = j
				tmp = arr[j]
			}
		}
		if k != i {
			arr[k], arr[i] = arr[i], arr[k]
		}
	}
	return arr
}

//插入排序 一个元素是有序的，将每个元素按照顺序插入有序元组中就是有序的
func InsertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			//找出要插入的位置
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

//快速排序 分成3部分 [左 一个基准点 右],一个元素是有序的
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	tmp := arr[0]
	left := []int{}
	right := []int{}

	for i := 1; i < len(arr); i++ {
		if arr[i] < tmp {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	//排序左边的
	left = QuickSort(left)
	left = append(left, tmp)
	//排序右边的
	right = QuickSort(right)
	return append(left, right...)
}
