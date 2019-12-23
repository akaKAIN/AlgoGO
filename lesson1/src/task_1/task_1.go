package task_1

import (
	"fmt"
	"log"
)

func SumTreeDigit() int {
	var input, sum int

	input = InterNum()
	sum = CalcSum(input)
	return sum
}

func CalcSum(num int)(sum int){
	if num < 0 {
		num *= -1
	}
	sum = num/100 + (num%100)/10 + (num % 10)
	return
}

func InterNum() (input int) {
	fmt.Print("Задание1.\nПодсчет суммы трехзначного числа\n")
InputLoop:
	fmt.Print("\nВведите трехзначное число\n")
	if _, err := fmt.Scan(&input); err != nil {
		log.Println("Error input ", err)
		goto InputLoop
	}
	row := fmt.Sprintf("%d", input)
	if len(row) != 3 {
		log.Println("Вы ввели число неверное по числу знаков.")
		goto InputLoop
	}
	return
}
