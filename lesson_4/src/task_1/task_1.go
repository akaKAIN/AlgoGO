package task_1

import (
	"fmt"
	"math/rand"
	"time"
)

func init () {
	rand.Seed(time.Now().UnixNano())
}

func Sieve () {
	var arr []int
	for i:= 1; i<1e6; i++{
		arr = append(arr, i)
	}
	fmt.Println(len(arr))
}