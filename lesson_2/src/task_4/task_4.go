package task_4

import (
	"fmt"
	"log"
)

/*"""
4.	Найти сумму n элементов следующего ряда чисел: 1 -0.5 0.25 -0.125 ...
Количество элементов (n) вводится с клавиатуры.
"""
 */

func HalfDivRow(){
	var input int
	var result float64
	result = 1

	if _, err := fmt.Scan(&input); err != nil {
		log.Printf("Ошибка ввода: %s, input=%v", err, input)
	}
	for i:=1; i<input; i++{
		result /= 2
	}
	fmt.Println(result)
}