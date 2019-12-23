package main

import "fmt"

func main() {
	var b = make([]byte, 10)
	for i:=0; i<10; i++{
		b[i] = 's'
	}
	fmt.Println(string(b))
}
