package task_8

import (
	"fmt"
	"log"
)

func Year(){
	var year int
	var input string
	YearLoop:
	fmt.Println("Для проверки года на высокосТность, введите номер года: ")
	if _, err := fmt.Scan(&year); err != nil {
		log.Printf("Ошибка ввода номера года: %s, вы ввели - %v", err, year)
		goto YearLoop
	}
	switch {
	case (year % 4 == 0 && year % 100 !=0) || year % 400 == 0:
		fmt.Printf("%d - это высокосТный год.", year)
	default:
		fmt.Print("Год самый обычный и не высокосный\n")
	}
	fmt.Println("\n\tХотете проверить другой год? (Y/N)")
	if _, err := fmt.Scan(&input); err != nil {
		log.Printf("Ошибка чтения ответа: %s, вы ввели - %v", err, input)
	}
	switch input {
	case "Y", "y", "Yes", "yes", "да", "ДА", "Да":
		goto YearLoop
	default:
		fmt.Println("Тогда, досвиданья.")
	}
}
