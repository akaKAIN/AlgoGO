package task_1

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

/*
1.	Написать программу, которая будет складывать, вычитать, умножать или делить
два числа. Числа и знак операции вводятся пользователем. После выполнения
вычисления программа не должна завершаться, а должна запрашивать новые данные
для вычислений. Завершение программы должно выполняться при вводе символа '0'
в качестве знака операции. Если пользователь вводит неверный знак
(не '0', '+', '-', '*', '/'), то программа должна сообщать ему об ошибке и
снова запрашивать знак операции.
Также сообщать пользователю о невозможности деления на ноль,
если он ввел 0 в качестве делителя.
*/

func Calc() {
	var row, symb string
	var result int64
	var choise int
	var err error

InputLoop:
	fmt.Println("\nВведите математическую операцию: ")
	if _, err := fmt.Scan(&row); err != nil {
		log.Printf("Error input: %s, input=%q", err, row)
		goto InputLoop
	}
	switch {
	case strings.Contains(row, "+"):
		symb = "+"
		result, err = parse(row, symb)
		if err != nil {
			fmt.Println(err)
			goto InputLoop
		}

	case strings.Contains(row, "-"):
		symb = "-"
		result, err = parse(row, symb)
		if err != nil {
			fmt.Println(err)
			goto InputLoop
		}
	case strings.Contains(row, "*"):
		symb = "*"
		result, err = parse(row, symb)
		if err != nil {
			fmt.Println(err)
			goto InputLoop
		}
	case strings.Contains(row, "/"):
		symb = "/"
		result, err = parse(row, symb)
		if err != nil {
			fmt.Println(err)
			goto InputLoop
		}
	case row == "0":
		fmt.Println("Инициирован выход из калькулятора.")
		time.Sleep(time.Second * 2)
		goto endLoop
	}
bufferLoop:
	fmt.Printf("Ответ: %d\n", result)
	fmt.Println("Желаете повторить?\n1-Да/2-Нет")
	if _, err := fmt.Scan(&choise); err != nil {
		log.Printf("Ошибка выбора %s", err)
		goto bufferLoop
	}
	switch choise {
	case 1: goto InputLoop
	case 2: goto endLoop
	default: goto bufferLoop
	}

endLoop:
	fmt.Println("bye!")

}

func parse(row string, operand string) (int64, error) {
	var ind int
	var res int64
	ind = strings.Index(row, operand)
	a, err := strconv.ParseInt(row[:ind], 10, 64)
	if err != nil {
		err = fmt.Errorf("Ошибка данных %s. input_row=%v, symb=%v\n", err, row, operand)
		return 0, err
	}
	b, err := strconv.ParseInt(row[:ind], 10, 64)
	if err != nil {
		err = fmt.Errorf("Ошибка данных %s. input_row=%v, symb=%v\n", err, row, operand)
		return 0, err
	}
	switch operand {
	case "+": res = a + b
	case "-": res = a - b
	case "*": res = a * b
	case "/": res = a / b
	}
	return res, nil
}
