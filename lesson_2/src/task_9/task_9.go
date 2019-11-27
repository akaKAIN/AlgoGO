package task_9

import (
	"fmt"
	"log"
)

/*
"""
9. Среди натуральных чисел, которые были введены, найти
наибольшее по сумме цифр. Вывести на экран это число и сумму его цифр.
"""
 */

func SumNumberUnit(){
	var quantity, num, sum, max, maxNum, baseNum int
	max = 0
	fmt.Println("Введите кол-во натуральных чисел для подсчета суммы их цифр: ")
	if _, err := fmt.Scan(&quantity); err != nil {
		log.Printf("Ошибка ввода %s, input=%v", err, quantity)
	}
	fmt.Printf("Введите %d натуральных чисел:\n", quantity)

	for i:=1; i<=quantity; i++{
		fmt.Print(i, " число = ")
		if _, err := fmt.Scan(&num); err != nil {
			log.Printf("Ошибка ввода числа: %s, введено:%v", err, num)
			fmt.Println("Введите число заново: \n")
			continue
		}
		baseNum = num
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		if sum > max {
			max = sum
			maxNum = baseNum
		}
		sum = 0
	}
	fmt.Printf("\nСамая большая сумма цифр: %d\nВ числе %d", max, maxNum)

}