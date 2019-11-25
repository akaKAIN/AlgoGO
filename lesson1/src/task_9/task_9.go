package task_9

import (
	"fmt"
	"log"
)

func MiddleNumber(){
	var a, b, c, middle int
	fmt.Println("Введите три числа: ")
	if _, err := fmt.Scan(&a, &b, &c); err != nil {
		log.Printf("Ошибка ввода данных %s: a, b, c = (%v, %v, %v)", err, a, b, c)
	}
	switch {
	case (a < b && b < c) || (a>b && b > c):
		middle = b
	case (c > a && a > b) || (c<a && a < b):
		middle = a
	default:
		middle = c
	}
	fmt.Printf("Из чисел: %d, %d, %d среднее - это %d", a, b, c, middle)
}
