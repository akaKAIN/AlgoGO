package task_1

import (
	"fmt"
	"log"
)

/*
# 1.	В диапазоне натуральных чисел от 2 до 99 определить,
# сколько из них кратны каждому из чисел в диапазоне от 2 до 9.
*/

func MultiplicationCount() {
	var limit int
	result := make(map[int]int)
	fmt.Println("Введите предел диапазона (с 2 до .. )")
	if _, err := fmt.Scan(&limit); err != nil {
		log.Printf("Ошибка ввода данных %s, %v", err, limit)
	}
	for i := 2; i <= limit; i++ {
		for j := 1; j <= 9; j++ {
			if i%j == 0 {
				result[j]++
			}
		}
	}
	for i := 1; i <= len(result); i++ {
		fmt.Printf("кратных %d: %d\n", i, result[i])
	}

}
