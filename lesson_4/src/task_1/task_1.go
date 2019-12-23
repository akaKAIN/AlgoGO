package task_1

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func PowFunc(num, pow int) int {
	if pow < 1 {
		return 1
	}
	return num * PowFunc(num, pow-1)
}


func genArray(lenArr int) []int {
	var obj []int
	limit := PowFunc(10, lenArr)
	for i := 1; i < limit; i++ {
		obj = append(obj, i)
	}
	return obj
}

func Sieve() {
	arr := genArray(6)
	var arr2 = make([]int, 10)
	copy(arr2, arr[:10])
	fmt.Println(arr2)

	arr3 := append([]int{}, arr[:20]...)
	fmt.Println(arr3)
}
