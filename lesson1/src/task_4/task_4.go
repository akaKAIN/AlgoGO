package task_4

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"unicode/utf8"
)


func MyRange(){
	rand.Seed(time.Now().UnixNano())
	fmt.Print("Написать программу, которая генерирует в указанных пользователем границах\n" +
	"●	случайное целое число,\n" +
	"●	случайное вещественное число,\n" +
	"●	случайный символ.\n" +
		"Для каждого из трех случаев пользователь задает свои границы диапазона.\n" +
		"Например, если надо получить случайный символ от 'a' до 'f',\n" +
		"то вводятся эти символы. Программа должна вывести на экран любой\n" +
	"символ алфавита от 'a' до 'f' включительно.\n\n")
Menu:
	fmt.Print("Выберите тип, желаемого диапазона:\n1. Случайное целое число.\n" +
		"2. Случайное вещественное число.\n3. случайный символ.\n\n>>>")
	var choise uint8
	_, err := fmt.Scan(&choise)
	if err != nil {
		log.Printf("Error choise: %s", err)
	}
	switch choise {
	case 1:
		NumberFunc()
	case 2:
		FloatFunc()
	case 3:
		SimbolFunc()
	default:
		fmt.Print("Неверный выбор пункта меню.")
		goto Menu
	}
}

func NumberFunc(){
IntLoop:
	fmt.Print("Введите через пробел верхний и нижние диапазоны\n")
	var min, max int
	_, err := fmt.Scan(&min, &max)
	if err != nil {
		log.Printf("Ошибка ввода int диапазона: %s", err)
		goto IntLoop
	}
	if min > max {
		min, max = max, min
	}
	numInRange := min + rand.Intn(max - min + 1)
	fmt.Printf("done. => [%d;%d]=%d\n", min, max, numInRange)
}

func FloatFunc(){
FloatLoop:
	fmt.Print("Введите через пробел верхний и нижние диапазоны\n")
	var min, max, minFloat, maxFloat, randFloat, Result float32
	var minInt, maxInt, randInt int
	_, err := fmt.Scan(&min, &max)
	if err != nil {
		log.Printf("Ошибка ввода float диапазона: %s", err)
		goto FloatLoop
	}
	if min > max {
		min, max = max, min
	}
	minInt, maxInt = int(min), int(max)
	minFloat, maxFloat = min - float32(minInt), max - float32(maxInt)
	randInt = minInt + rand.Intn(maxInt - minInt + 1)
RandomFloatLoop:
	if minFloat > maxFloat {
			randFloat = rand.Float32()
			if randFloat > ((maxFloat - minFloat)*(-1)){
				goto RandomFloatLoop
			}

	}else{
		randFloat = rand.Float32()
		if randFloat < minFloat || randFloat > maxFloat {
			goto RandomFloatLoop
		}
	}
	Result = randFloat + float32(randInt)
	fmt.Println(Result)
}

func SimbolFunc(){
	var min, max string
	if _, err := fmt.Scan(&min, &max); err != nil{
		log.Printf("Error of input string %s", err)
	}

	runeMin, _:=utf8.DecodeRuneInString(min)
	runeMax, _ :=utf8.DecodeLastRuneInString(max)
	if int(runeMin) > int(runeMax) {
		runeMin, runeMax = runeMax, runeMin
	}
	randRuneInt := int(runeMin) + rand.Intn(int(runeMax) - int(runeMin) + 1)
	str := fmt.Sprint(rune(randRuneInt))
	fmt.Printf( "%#U", str)




}
