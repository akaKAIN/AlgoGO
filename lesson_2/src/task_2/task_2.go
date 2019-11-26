package task_2

import (
	"fmt"
	"log"
)

/*
"""
2.	Посчитать четные и нечетные цифры введенного натурального числа.
Например, если введено число 34560, то у него 3 четные цифры
(4, 6 и 0) и 2 нечетные (3 и 5).
"""
 */

func CountSameNumbers(){
	var err error
	var num, even, odd, unit int

InputLoop:
	fmt.Println("Введите натуральное число:")
	if _, err = fmt.Scan(&num); err != nil {
		log.Printf("Ошибка ввода данных: %s", err)
		goto InputLoop
	}
	for num >0 {
		unit = num % 10
		num /= 10
		if unit % 2 == 0 {
			even++
		}else{
			odd++
		}
	}
	fmt.Printf("Even=%d  odd=%d", even, odd)
}