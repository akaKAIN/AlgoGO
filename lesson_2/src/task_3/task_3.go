package task_3

import (
	"fmt"
	"log"
)

/*"""
3.	Сформировать из введенного числа обратное по порядку входящих в него
цифр и вывести на экран. Например, если введено число 3486,
 то надо вывести число 6843.
"""
 */

func ReversPrint(){
	var input int
	fmt.Print("Введите число, для вывода его отзеркаленной версии.")
	if _, err := fmt.Scan(&input); err != nil {
		log.Printf("Ошибка ввода данных: %s;\ninput=%v", err, input)
	}
	for input > 0 {
		fmt.Print(input%10)
		input /= 10
	}

}