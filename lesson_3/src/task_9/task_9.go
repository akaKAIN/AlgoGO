package task_9

import "fmt"

/*
Найти максимальный элемент среди минимальных элементов столбцов матрицы.
*/
func ColumnSum() {
	var matrix = [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 1, 14, 1},
		{16, 1, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	var result = [5]int{1e4, 1e4, 1e4, 1e4, 1e4,}
	var max int

	for ind, _ := range matrix {

		for key, val := range matrix[ind] {
			if val < result[key] {
				result[key] = val
			}
		}
	}
	for _, val := range result {
		if val > max {
			max = val
		}
	}
	fmt.Println("Максимальное число их минимальных чисел каждой колонки матрицы:", max)
}
