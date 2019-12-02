package task_8

import (
	"fmt"
	"log"
)

/*
Матрица 5x4 заполняется вводом с клавиатуры кроме последних элементов строк.
Программа должна вычислять сумму введенных элементов каждой строки и
записывать ее в последнюю ячейку строки.
В конце следует вывести полученную матрицу.
*/

type Matrix [5][4]int

func MatrixSum() {
	fmt.Println("Введите числа, матрицы 5х4: ")
	matrix := new(Matrix)
	for i := 0; i < len(matrix); i++ {
		InputRow(&matrix[i])
		RowSum(&matrix[i])
	}
	fmt.Println(*matrix)
}
func InputRow(row *[4]int) {
	if _, err := fmt.Scan(&row[0], &row[1], &row[2]); err != nil {
		log.Println("Error: ", err)
	}
}

func RowSum(row *[4]int) {
	for i := 0; i < len(row)-1; i++ {
		row[3] += row[i]
	}
}
