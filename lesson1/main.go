package main

import (
	"fmt"
	"log"
	"time"
)

func main(){
	//fmt.Printf("Сумма=%d\n\n", task_1.SumTreeDigit())
	//task_2.BoolOperations()
	//task_3.TwoPoints()
	//task_3.TwoPoints()
	type inp interface {}
	//var n inp
	var name int
StartLoop:
	time.Sleep(2*time.Second)
	if _, err := fmt.Scan(&name); err != nil {
		log.Printf("Ошибка ввода данных: %s", err)
		goto StartLoop
	}
	fmt.Printf("Type=%T, val=%v", name, name)

}
