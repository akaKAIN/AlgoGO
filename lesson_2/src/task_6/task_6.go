package task_6

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

/*"""
6.	В программе генерируется случайное целое число от 0 до 100.
Пользователь должен его отгадать не более чем за 10 попыток. После каждой
неудачной попытки должно сообщаться больше или меньше введенное пользователем
число, чем то, что загадано. Если за 10 попыток число не отгадано,
то вывести загаданное число.
"""
 */

func Vanga(){
	rand.Seed(time.Now().UnixNano())
	var num, choice, count int
	num = rand.Intn(100)
GesLoop:
	fmt.Println("введите номер предполагаемого числа:")
	if _, err := fmt.Scan(&choice); err != nil {
		log.Printf("Ошибка ввода: %s - %v", err, choice)
	}
	switch {
	case count == 10:
		fmt.Printf("Вы ПРОИГРАЛИ.")
		return
	case choice < num:
		fmt.Printf("Число  %d - меньше загаданного.", choice)
		count++
		goto GesLoop
	case choice > num:
		fmt.Printf("Число  %d - больше загаданного.", choice)
		count++
		goto GesLoop
	case choice == num:
		fmt.Println("ПОБЕДА! Вы угадали число", num)
	}

}