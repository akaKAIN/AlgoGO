package task_1

import (
	"fmt"
	"log"
)

func SumTreeDigit() int {
	var input, sum int
	fmt.Print("Задание1.\nПодсчет суммы трехзначного числа\nВведите трехзначное число\n")
	if _, err := fmt.Scan(&input); err != nil {
		log.Println("Error input ", err)
	}
	if input < 0 {
		input *= -1
	}
	sum = input/100 + (input % 100)/10 + (input % 10)
	return sum
}
