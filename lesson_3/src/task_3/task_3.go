package task_3

import (
	"fmt"
	"task_2"
)

/*
#3.	В массиве случайных целых чисел поменять местами минимальный
# и максимальный элементы.
 */

func ChangeMinMax(){
	var randArr []int
	var min, max int = 0, 0
	randArr = task_2.GenArray()
	for _, val := range randArr {
		if val < min {
			min = val
		}else if val > max {
			max = val
		}
	}
	fmt.Println(randArr)
	fmt.Printf("Min=%d, Max=%d\n", min, max)
	for ind, val := range randArr {
		if val == min {
			randArr[ind] = max
		}else if val == max {
			randArr[ind] = min
		}
	}
	fmt.Println(randArr)
}
