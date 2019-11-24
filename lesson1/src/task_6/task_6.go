package task_6

import (
	"fmt"
	"log"
	"unicode/utf8"
)

func FindSymbol(){
	var input int32
	fmt.Println("Введите номер буквы алфавита:")
	if _, err := fmt.Scan(&input); err != nil {
		log.Printf("Error input: %s, input=%v", err, input)
	}
	symbol, _ := utf8.DecodeRuneInString(string(input + 96))
	fmt.Printf("Введеной цифре %d соответствует буква %q", input, symbol)
}