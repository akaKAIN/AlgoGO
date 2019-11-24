package task_7

import (
	"fmt"
	"log"
)

func Triangle() {
	var a, b, c int
	var triangleType string
InputLoop:
	fmt.Println("Введите длины сторон треугольника: ")
	if _, err := fmt.Scan(&a, &b, &c); err != nil {
		log.Printf("Error data input %s, a,b,c=(%v, %v, %v)", err, a, b, c)
		goto InputLoop
	}
	if checkTriangle(a, b, c){
		switch {
		case a == b && b == c:
			triangleType = "равносторонний"
		case a == b || b == c || c == a:
			triangleType = "равнобедренный"
		case a != b && b != c:
			triangleType = "разносторонний"
		default:
			triangleType = "существует и ладно, "
		}
		fmt.Printf("Введеные Вами стороны треугольника, являются сторонами треугольника, который %s", triangleType)
	}else{
		fmt.Println("Треугольник с такими сторонами не существует")
	}


}

//Проверка существования треуольника
func checkTriangle(a, b, c int) bool {
	if a+b > c && a+c > b && b+c > a {
		return true
	}else {
		return false
	}
}
