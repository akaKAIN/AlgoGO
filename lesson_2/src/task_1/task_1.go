package task_1

import (
	"fmt"
	"log"
	"strings"
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
	var result int

	InputLoop:
	if _, err := fmt.Scan(&row); err != nil {
		log.Printf("Error input: %s, input=%q", err, row)
		goto InputLoop
	}
	switch {
	case strings.Contains(row, "+"):
		symb = "+"
		result = add(parse(row, symb))
	case strings.Contains(row, "-"):
		sub(row)
	case strings.Contains(row, "*"):
		milt(row)
	case strings.Contains(row, "/"):
		div(row)
	}
}

func parse(row string, operand string) (a, b){
	a := 0
	b:= 0
	return
}

func add(row) {

}
