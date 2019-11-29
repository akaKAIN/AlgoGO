package task_6

import (
	"fmt"
)

/*
В одномерном массиве найти сумму элементов, находящихся
между минимальным и максимальным элементами.
Сами минимальный и максимальный элементы в сумму не включать.
*/
type Point struct {
	ind int
	val int
}

type Result struct {
	Min Point
	Max Point
	Sum int
}

func RangeSum() {
	var result Result
	var start, end int
	//randArr := task_2.GenArray()
	randArr := []int{1, 2, 3, -3, 20, 5, 10, -20, 15}
	for ind, val := range randArr {
		if val > result.Max.val {
			result.Max = Point{ind, val}
		}
		if val < result.Min.val {
			result.Min = Point{ind, val}
		}
	}

	switch {
	case result.Min.ind < result.Max.ind:
		start = result.Min.ind
		end = result.Max.ind
	default:
		start = result.Max.ind
		end = result.Min.ind
	}
	for i := start + 1; i < end; i++ {
		result.Sum += randArr[i]
	}

	//fmt.Println(randArr)
	fmt.Printf("Сумма чисел в массиве между %d  и  %d  равна %d\n", randArr[start], randArr[end], result.Sum)
	fmt.Println(result)

}
