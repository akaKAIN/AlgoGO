package task_7

import (
	"fmt"
	"log"
)

func Triangle(){
	var a, b, c int
	fmt.Println("Введите длины сторон треугольника: ")
	if _, err := fmt.Scan(&a, &b, &b); err != nil {
		log.Printf("Error data input %s, a,b,c=%v%v%v", err, a,b,c)
	}

}