package task_7

import (
	"fmt"
)

/*
В одномерном массиве целых чисел определить два наименьших элемента.
Они могут быть как равны между собой (оба являться минимальными),
 так и различаться.
*/
func MinTwoNum() {
	var first, second int

	randArr := []int{1, 2, 3, 4, 5, 6, -4, 7, 7, 8, 8, 9, -2}
	first, second = randArr[0], randArr[1]
	for _, val := range randArr {
		if val < second || val < first {
			first, second = second, val
		}
	}
	fmt.Println(randArr, first, second)
}
