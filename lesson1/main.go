package main

import (
	"fmt"
	"task_1"
	"task_2"
	"task_3"
	"task_4"
	"task_5"
	"task_6"
	"task_7"
	"task_8"
	"task_9"

)

func main(){
	fmt.Printf("Сумма=%d\n\n", task_1.SumTreeDigit())
	task_2.BoolOperations()
	task_3.TwoPoints()
	task_4.MyRange()
	if err := task_5.FindRangeSymbol(); err != nil {
		fmt.Println(err)
	}
	task_6.FindSymbol()
	task_7.Triangle()
	task_8.Year()
	task_9.MiddleNumber()
}
