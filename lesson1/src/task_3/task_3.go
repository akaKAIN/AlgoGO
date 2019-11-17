package task_3

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	Ox int32
	Oy int32
}

func TwoPoints() {
	fmt.Print("Задание 3.\nПо введенным пользователем координатам двух точек вывести\n " +
		"уравнение прямой вида y = kx + b, проходящей через эти точки.")

	var buf string
StartA:
	fmt.Println("Введите через '-' координаты точки А\n")
	if _, err := fmt.Scan(&buf); err != nil {
		fmt.Println("Error- ", err)
		goto StartA
	}
	xA := buf[:strings.Index(buf, "-")]
	yA := buf[strings.Index(buf, "-")+ 1:]
	xAInt, err := strconv.ParseInt(xA, 10, 32)
	if err != nil {
		fmt.Println("Error of parse int: ", err)
		goto StartA
	}
	yAInt, err := strconv.ParseInt(yA, 10, 32)
	if err != nil {
		fmt.Println("Error of parse int: ", err)
		goto StartA
	}
	pointA := Point{int32(xAInt), int32(yAInt)}
StartB:
	fmt.Println("Введите через '-' координаты точки B\n")
	if _, err := fmt.Scan(&buf); err != nil {
		fmt.Println("Error- ", err)
		goto StartB
	}
	xB := buf[:strings.Index(buf, "-")]
	yB := buf[strings.Index(buf, "-")+1:]
	xBInt, err := strconv.ParseInt(xB, 10, 32)
	if err != nil {
		fmt.Println("Error of parse int: ", err)
		goto StartB
	}
	yBInt, err := strconv.ParseInt(yB, 10, 32)
	if err != nil {
		fmt.Println("Error of parse int: ", err)
		goto StartB
	}
	pointB := Point{int32(xBInt), int32(yBInt)}

	//Вычисляем формулу
	k := (pointB.Oy - pointA.Oy) / (pointB.Ox - pointA.Ox)
	b := pointA.Oy - (pointB.Oy-pointA.Oy)*pointA.Ox/(pointB.Ox-pointA.Ox)
	fmt.Printf("Вам нужно уравнение вида:\ny = %dx + %d", k, b)
}
