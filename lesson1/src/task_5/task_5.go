package task_5

import (
	"fmt"
	"unicode/utf8"
)

func FindRangeSymbol() error {
	fmt.Printf("Пользователь вводит две буквы.\nОпределить, на каких места алфавита они стоят,\nи сколько между ними находится букв\n")
	var start, end string
	if _, err := fmt.Scan(&start, &end); err != nil {
		err = fmt.Errorf("Ошибка ввода:  %s, %s, %s", start, end, err)
		return err
	}
	startByte, _ := utf8.DecodeRune([]byte(start))
	endByte, _ := utf8.DecodeRune([]byte(end))
	fmt.Printf("Позиции введенных букв:\n%q на %d\n%q на %d\nИ между ними букв: %d", start, startByte - 96, end, endByte-96, endByte - startByte)
	return nil
}
