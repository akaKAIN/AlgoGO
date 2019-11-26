package task_8

import (
	"fmt"
	"log"
)

/*
"""
8.	Посчитать, сколько раз встречается определенная цифра в введенной
 последовательности чисел. Количество вводимых чисел и цифра,
 которую необходимо посчитать, задаются вводом с клавиатуры.
"""
 */

func Counter() {
	var(
		input, sym string
		counter int
	)
	fmt.Println("Введите любое число (или строку)")
	if _, err := fmt.Scan(&input); err != nil {
		log.Printf("Ошибка ввода данных: %s, %v", err, input)
	}
	fmt.Print("Введите символ для подсчета - ")
	if _, err := fmt.Scan(&sym); err != nil {
		log.Printf("Ошибка ввода данных: %s, %v", err, input)
	}
	for _, elem := range input {
		if string(elem) == sym {
			counter++
		}
	}
	fmt.Printf("В введеном %q найдено: %d повторений %q", input, counter, sym)
}