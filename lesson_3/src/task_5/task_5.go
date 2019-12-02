package task_5

import (
	"fmt"
)

/*
В массиве найти максимальный отрицательный элемент.
# Вывести на экран его значение и позицию (индекс) в массиве.
*/

type Point struct {
	Ind int
	Val int
}

func MaxNegativeNumber() {
	var num Point
	randArr := []int{1, 2, -3, -4, -5, 10, -8, -44}
	for key, val := range randArr {
		if val < 0 && val < num.Val {
			num = Point{key, val}
		}
	}
	if num.Val >= 0 {
		fmt.Println("Отлицательныз чисел не найдено в массиве.")
	} else {
		fmt.Printf("Самое большое отрицательное = %d по индексу %d\n", num.Val, num.Ind)
	}
	fmt.Println(randArr)

}
