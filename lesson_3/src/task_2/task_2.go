package task_2

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

/*
2.	Во втором массиве сохранить индексы четных элементов первого массива.
Например, если дан массив со значениями 8, 3, 15, 6, 4, 2, то во второй массив
надо заполнить значениями 1, 4, 5, 6
(или 0, 3, 4, 5 - если индексация начинается с нуля),
т.к. именно в этих позициях первого массива стоят четные числа.
"""
 */

func genArray(limit int) []int {
	result := make([]int, 0)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<limit; i++{
		result = append(result, rand.Intn(limit))
	}
	return result
}

func SortArray(){
	var limit int
	var randArr, sortArr []int


	fmt.Println("Введите длину массива")
	if _, err := fmt.Scan(&limit); err != nil {
		log.Printf("Ошибка ввода %s, %v", err, limit)
	}
	randArr = genArray(limit)
	for ind, val := range randArr {
		if val % 2 == 0 {
			sortArr = append(sortArr, ind)
		}
	}
	fmt.Println("Сгенерированный массив: ",randArr)
	fmt.Println("Отсеянный массив: ", sortArr)
}


