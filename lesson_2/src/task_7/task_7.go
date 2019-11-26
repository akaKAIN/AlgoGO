package task_7

import (
	"fmt"
	"log"
)

/*"""
7.	Напишите программу, доказывающую или проверяющую, что для множества
натуральных чисел выполняется равенство: 1+2+...+n = n(n+1)/2,
 где n - любое натуральное число.
"""
*/

func Pifag() {
	var input, sum, pif int
	if _, err := fmt.Scan(&input); err != nil {
		log.Printf("Ошибка %s, %v", err, input)
	}
	for i := 0; i <= input; i++ {
		sum += i
	}
	pif = input*(input+1)/2
	fmt.Println("Сумма чисел 1 + 2 + ... + n =", sum)
	fmt.Println("Выражение вида n * (n + 1) / 2 =", pif)
	if sum == pif {
		fmt.Println("Равенство доказано.")
	}else{
		fmt.Println("Странно, но они оказались неравны ...")
	}
}
